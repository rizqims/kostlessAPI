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
	AccPayment(payload dto.AccPayment) (string, error)
}

type transService struct {
	transRepo  repository.TransRepo
	userRepo   repository.UserRepo
	seekerRepo repository.SeekerRepo
	roomRepo   repository.RoomRepository
	voucherRepo repository.VoucherRepo
}

func (t *transService) CreateTrans(payload dto.TransCreateReq) (model.Trans, error) {
	var trans model.Trans
	seeker, err := t.seekerRepo.GetSeekerByID(payload.SeekerID)
	if err != nil {
		return model.Trans{}, fmt.Errorf("error in GetSeekerByID: %v", err)
	}

	startDate, err := time.Parse(`2006-01-02`, payload.StartDate)
	if err != nil {
		return model.Trans{}, fmt.Errorf("error in time.Parse: %v", err)
	}
	trans.EndDate = startDate.AddDate(0, payload.Months, 0)

	// implement total logic
	room, err := t.roomRepo.GetRoomByID(payload.RoomID)
	if err != nil {
		return model.Trans{}, fmt.Errorf("error in GetRoomByID: %v", err)
	}

	totalPrice := room.Price * payload.Months

	// validate date
	if trans.EndDate.Before(startDate) {
		return model.Trans{}, errors.New("endDate should not before StartDate")
	} else if startDate.Before(time.Now()){
		return model.Trans{}, errors.New("startDate should be in the future")
	}

	//validate voucher
	if payload.VoucherID == ""{
		trans.Discount = 0
		trans.Total = totalPrice
		fmt.Print("this get printed")
	} else {
		if payload.Months != 1{
			return model.Trans{}, fmt.Errorf("error in checkMonth: cannot rent over than 1 month if using voucher")
		}
		voucher, err := t.voucherRepo.GetVoucherByID(payload.VoucherID)
		if err != nil {
			return model.Trans{}, fmt.Errorf("error in getvoucher: %v", err)
		} else if voucher.SeekerID != payload.SeekerID{
			return model.Trans{}, fmt.Errorf("error in seekermatch : %v", err)
		}

		// pusing pala gwej
		discountedPrice := (float32(voucher.PercentAmount) / 100) * float32(totalPrice)
		trans.Discount = voucher.PercentAmount
		trans.Total = int(discountedPrice)
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
		trans.PaymentStatus = "pending"
		falseDueDate, err := time.Parse(`2006-01-02`, `2024-01-01`)
		if err != nil {
			return model.Trans{}, err
		}
		trans.DueDate = falseDueDate
	}

	trans.RoomID = payload.RoomID
	trans.SeekerID = payload.SeekerID
	trans.StartDate = startDate
	trans.PayLater = payload.PayLater

	transReq, err := t.transRepo.CreateTrans(trans)
	if err != nil {
		return model.Trans{}, fmt.Errorf("createtransrepo error: %v", err)
	}

	// TODO use detailed message
	err = notifyTransToUsers(
		seeker.Email,
		"New Booking Request",
		fmt.Sprintf("seeker named %v wants to books your Kost", seeker.Fullname),
		fmt.Sprintf("localhost:8080/api/v1/trans/"),
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

	seeker, err := t.seekerRepo.GetSeekerByID(payload.SeekerID)
	if err != nil {
		return dto.UpdatePaylaterRes{}, err
	}

	err = notifyTransToUsers(
		"rizqims100@gmail.com",
		"Paylater payment request",
		fmt.Sprintf("seeker named %v wants to pay previous pending payment", seeker.Fullname),
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

func (t *transService) AccPayment(payload dto.AccPayment) (string, error){	
	result, err := t.transRepo.AccPayment(payload)
	if err != nil {
		return "", err
	}

	seeker, err := t.seekerRepo.GetSeekerByID(result)
	if err != nil {
		return "", err
	}

	notifyTransToUsers(
		seeker.Email,
		"Transaction payment",
		"Your transaction has been validated by the owner. Below is the details of your transaction",
		fmt.Sprintf("localhost:8080/api/v1/trans/%v", payload.TransID),
	)

	return "success accepting payment, seeker will be notified", nil
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

func NewTransService(transRepo repository.TransRepo, userRepo repository.UserRepo, seekerRepo repository.SeekerRepo, roomrepo repository.RoomRepository, voucherRepo repository.VoucherRepo) TransService {
	return &transService{
		transRepo:  transRepo,
		userRepo:   userRepo,
		seekerRepo: seekerRepo,
		roomRepo:   roomrepo,
		voucherRepo: voucherRepo,
	}
}
