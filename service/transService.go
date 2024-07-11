package service

import (
	"errors"
	"fmt"
	"kostless/model"
	"kostless/model/dto"
	"kostless/repository"
	"net/smtp"
	"strconv"
	"time"
)

type TransService interface {
	CreateTrans(payload dto.TransCreateReq) (model.Trans, error)
	GetTransByID(id string) (model.Trans, error)
	GetTransHistory() ([]model.Trans, error)
	GetPaylaterList() ([]model.Trans, error)
	GetTransByMonth(month, year string) ([]model.Trans, error)
	UpdatePaylater(payload dto.UpdatePaylaterReq) (dto.UpdatePaylaterRes, error)
}

type transService struct {
	transRepo  repository.TransRepo
	userRepo   repository.UserRepo
	seekerRepo repository.SeekerRepo
	roomRepo   repository.RoomRepository
}

func (t *transService) CreateTrans(payload dto.TransCreateReq) (model.Trans, error) {
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
	err = notifyTransToUsers(
		"rizqims100@gmail.com",
		"New Booking Request",
		"seeker named Asep wants to books your Kost",
		"localhost:8080/api/v1/getall",
	)
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

func (t *transService) GetPaylaterList() ([]model.Trans, error) {
	transList, err := t.transRepo.GetPaylaterList()
	if err != nil {
		return nil, fmt.Errorf("GetPaylaterListService: get trans error: ", err)
	}
	return transList, nil
}

func (t *transService) GetTransByMonth(month, year string) ([]model.Trans, error) {
	rawDate := fmt.Sprintf("%v-%v-01", year, month)

	startDate, err := time.Parse(`2006-January-02`, rawDate)
	if err != nil {
		return nil, err
	}
	endDate := startDate.AddDate(0, 1, 0)
	fmt.Println(startDate, endDate)

	resStartDate := startDate.Format(`2006-01-02`)
	resEndDate := endDate.Format(`2006-01-02`)
	fmt.Println(resStartDate, resEndDate)

	transList, err := t.transRepo.GetTransByMonth(resStartDate, resEndDate)
	if err != nil {
		return nil, err
	}

	return transList, nil
}

func (t *transService) UpdatePaylater(payload dto.UpdatePaylaterReq) (dto.UpdatePaylaterRes, error) {
	trans, err := t.transRepo.UpdatePaylater(payload)
	if err != nil {
		return dto.UpdatePaylaterRes{}, err
	}
	currentDate := time.Now()
	var IsOverdue bool
	penaltyDays := currentDate.Format(`02`)
	converted, err := strconv.Atoi(penaltyDays)
	if err != nil {
		fmt.Println(err)
	}
	convFloat := float32(converted)
	penaltyPercent := (convFloat / 100) * 100
	totalWithPenalty := float32(trans.Total) / penaltyPercent

	if payload.Total != int(totalWithPenalty) {
		IsOverdue = true
		return dto.UpdatePaylaterRes{
			TransID:          payload.TransID,
			IsOverdue:        IsOverdue,
			DueDate:          trans.DueDate,
			CurrentTime:      currentDate,
			TotalWithPenalty: int(totalWithPenalty),
		}, errors.New("money is not enough to pay the transaction")
	}

	// seeker, err := t.seekerRepo.GetByID(payload.SeekerID)
	// if err != nil {
	// 	return model.Trans{}, err
	// }

	err = notifyTransToUsers(
		"rizqims100@gmail.com",
		"Paylater payment request",
		"seeker named Asep wants to pay previous pending payment",
		"localhost:8080/api/v1/paylaterlist",
	)
	if err != nil {
		return dto.UpdatePaylaterRes{}, err
	}

	return dto.UpdatePaylaterRes{
		TransID:          payload.TransID,
		IsOverdue:        IsOverdue,
		DueDate:          trans.DueDate,
		CurrentTime:      time.Now(),
		TotalWithPenalty: int(totalWithPenalty),
	}, nil
}

func notifyTransToUsers(to, subject, desc, link string) error {
	// Your AnonAddy SMTP credentials
	smtpHost := "smtp.mail.yahoo.com"
	smtpPort := "587"
	smtpUsername := "kumarpakcik@yahoo.com"
	smtpPassword := "wdtysyesezlfcbvn"

	from := "kumarpakcik@yahoo.com"
	body := fmt.Sprintf(`
		<html>
		<body>
		<h1>Welcome to Kostless</h1>
		<p>%v</p>
		<p>Click the link below to proceed:</p>
		<a href="%v">here</a>
		</body>
		</html>
		`, desc, link)

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
