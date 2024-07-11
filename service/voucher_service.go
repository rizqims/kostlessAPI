package service

import (
	"kostless/model"
	"kostless/model/dto"
	"kostless/repository"
)

type VoucherService interface {
	CreateVoucher(payload dto.CreateVoucherReq) (model.Voucher, error)
	DeleteExpiredVoucher() error
	GetAllVoucher() ([]model.Voucher, error)
	GetVoucherBySeekerID(id string) ([]model.Voucher, error)
}

type voucherService struct {
	repo repository.VoucherRepo
}

func (v *voucherService) CreateVoucher(payload dto.CreateVoucherReq) (model.Voucher, error) {
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


func NewVoucherService(repo repository.VoucherRepo) VoucherService {
	return &voucherService{
		repo: repo,
	}
}
