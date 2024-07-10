package service

import (
	"errors"
	"fmt"
	"kostless/model"
	"kostless/model/req"
	"kostless/repository"
	"net/smtp"
	"time"
)

type TransService interface {
	CreateTrans(payload req.TransCreateReq) (model.Trans, error)
	GetTransByID(id string) (model.Trans, error)
	GetTransHistory() ([]model.Trans, error)
	GetPaylaterList() ([]model.Trans, error)
}

type transService struct {
	transRepo  repository.TransRepo
	userRepo   repository.UserRepo
	seekerRepo repository.SeekerRepo
	roomRepo   repository.RoomRepository
}

func (t *transService) CreateTrans(payload req.TransCreateReq) (model.Trans, error) {
	var trans model.Trans

	// TODO kaitkan dengan user asli
	// user, err := t.userRepo.GetByUsername("aseppp")
	// if err != nil {
	// 	return model.Trans{}, err
	// }

	// seeker, err := t.seekerRepo.GetByID(payload.SeekerID)
	// if err != nil {
	// 	return model.Trans{}, err
	// }

	startDate, err := time.Parse(`2006-01-02`, payload.StartDate)
	if err != nil {
		return model.Trans{}, err
	}
	trans.EndDate = startDate.AddDate(0, payload.Months, 0)
	// TODO implement total logic
	// total, err := t.roomRepo.GetRoomByID(payload.RoomID)

	// validate date
	if trans.EndDate.Before(startDate) {
		return model.Trans{}, errors.New("endDate should not before StartDate")
	}

	// validate paylater
	if payload.PayLater {
		trans.PaymentStatus = "pending"
		trueDueDate, err := time.Parse(`2006-01-02`, payload.DueDate)
		if err != nil {
			return model.Trans{}, err
		}
		trans.DueDate = trueDueDate
	} else {
		trans.PaymentStatus = "paid"
		falseDueDate, err := time.Parse(`2006-01-02`, `2024-01-01`)
		if err != nil {
			return model.Trans{}, err
		}
		trans.DueDate = falseDueDate
	}

	trans.RoomID = payload.RoomID
	trans.SeekerID = payload.SeekerID
	trans.StartDate = startDate
	trans.Total = 10000 // hardcoded bcs no room yet
	trans.PayLater = payload.PayLater

	transReq, err := t.transRepo.CreateTrans(trans)
	if err != nil {
		return model.Trans{}, err
	}

	// TODO use detailed message
	err = notifyTransToUsers("rizqims100@gmail.com", model.Seekers{})
	if err != nil {
		return model.Trans{}, err
	}
	return transReq, nil
}

func (t *transService) GetTransByID(id string) (model.Trans, error) {
	if id == "" {
		return model.Trans{}, errors.New("id must be present")
	}

	trans, err := t.transRepo.GetTransByID(id)
	if err != nil {
		return model.Trans{}, nil
	}

	return trans, nil
}

func (t *transService) GetTransHistory() ([]model.Trans, error) {
	transList, err := t.transRepo.GetTransHistory()
	if err != nil {
		return nil, fmt.Errorf("GetTransHistoryService: get trans error: ", err)
	}
	return transList, nil
}

func (t *transService) GetPaylaterList() ([]model.Trans, error){
	transList, err := t.transRepo.GetPaylaterList()
	if err != nil {
		return nil, fmt.Errorf("GetPaylaterListService: get trans error: ", err)
	}
	return transList, nil
}

func notifyTransToUsers(to string, seeker model.Seekers) error {
	// Your AnonAddy SMTP credentials
	smtpHost := "smtp.mail.yahoo.com"
	smtpPort := "587"
	smtpUsername := "kumarpakcik@yahoo.com"
	smtpPassword := "wdtysyesezlfcbvn"

	from := "kumarpakcik@yahoo.com"
	subject := "New Booking Request"
	body := fmt.Sprintf(`
		<html>
		<body>
		<h1>Welcome to Kostless</h1>
		<p>seeker named %v wants to books your Kost</p>
		<p>Click the link below to proceed:</p>
		<a href="localhost:8080/api/v1/getall">here</a>
		</body>
		</html>
		`, "asep")

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func NewTransService(transRepo repository.TransRepo, userRepo repository.UserRepo, seekerRepo repository.SeekerRepo, roomrepo repository.RoomRepository) TransService {
	return &transService{
		transRepo:  transRepo,
		userRepo:   userRepo,
		seekerRepo: seekerRepo,
		roomRepo:   roomrepo,
	}
}
