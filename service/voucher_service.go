package service

import (
	"errors"
	"kostless/model"
	"kostless/model/dto"
	"kostless/repository"
	"time"
)

type VoucherService interface {
	CreateVoucher(payload dto.CreateVoucherReq) (model.Voucher, error)
	DeleteExpiredVoucher() error
	GetAllVoucher() ([]model.Voucher, error)
	GetVoucherBySeekerID(id string) ([]model.Voucher, error)
	GetVoucherByID(id string)(model.Voucher, error)
}

type voucherService struct {
	repo repository.VoucherRepo
}

func (v *voucherService) CreateVoucher(payload dto.CreateVoucherReq) (model.Voucher, error) {
	expired, err := time.Parse(`2006-01-02`, payload.ExpiredDate)
	if expired.Before(time.Now()){
		return model.Voucher{}, errors.New("expired date should not be in the past")
	}
	voucher, err := v.repo.CreateVoucher(payload)
	if err != nil {
		return model.Voucher{}, err
	}
	return voucher, nil
}

func (v *voucherService) DeleteExpiredVoucher() error {
	err := v.repo.DeleteExpiredVoucher()
	if err != nil {
		return err
	}
	return nil
}

func (v *voucherService) GetAllVoucher() ([]model.Voucher, error){
	voucherList, err := v.repo.GetAllVoucher()
	if err != nil {
		return nil, err
	}
	return voucherList, nil
}

func (v *voucherService) GetVoucherBySeekerID(id string) ([]model.Voucher, error){
	voucherList, err := v.repo.GetVoucherBySeekerID(id)
	if err != nil {
		return nil, err
	}
	return voucherList, nil
}

func (v *voucherService) GetVoucherByID(id string)(model.Voucher, error){
	voucherList, err := v.repo.GetVoucherByID(id)
	if err != nil {
		return model.Voucher{}, err
	}
	return voucherList, nil
}


func NewVoucherService(repo repository.VoucherRepo) VoucherService {
	return &voucherService{
		repo: repo,
	}
}
