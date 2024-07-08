package service

import (
	"errors"
	"kostless/model"
	"kostless/model/req"
	"kostless/repository"
	"time"
)

type TransService interface {
	CreateTrans(payload req.TransCreateReq) (model.Trans, error)
}

type transService struct {
  repo repository.TransRepo
}

func (t *transService) CreateTrans(payload req.TransCreateReq) (model.Trans, error){
	var trans model.Trans
	trans.EndDate = payload.StartDate.AddDate(0, payload.Months, 0)
	// TODO implement total logic
	// total := repository.getRoomPrice

	// validate date
	if trans.EndDate.Before(payload.StartDate){
		return model.Trans{}, errors.New("endDate should not before StartDate")
	}

	// validate paylater
	if payload.PayLater{
		trans.PaymentStatus = "pending"
	} else{
		trans.PaymentStatus = "paid"
		newDueDate, err := time.Parse(`2006-01-02`, `2024-01-01`)
		if err != nil {
			return model.Trans{}, err
		}
		trans.DueDate = newDueDate
	}


	trans.RoomID = payload.RoomID
	trans.SeekerID = payload.SeekerID
	trans.StartDate = payload.StartDate
	trans.Total = 10000
	trans.PayLater = payload.PayLater
	trans.DueDate = payload.DueDate

	transReq, err := t.repo.CreateTrans(trans)
	if err != nil {
		return model.Trans{}, err
	}
	return transReq, nil
}

func NewTransService(repo repository.TransRepo) TransService{
	return &transService{
		repo: repo,
	}
}