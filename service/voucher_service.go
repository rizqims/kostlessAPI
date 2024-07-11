package service

import (
	"kostless/model"
	"kostless/model/dto"
	"kostless/repository"
)

type VoucherService interface {
	CreateVoucher(payload dto.CreateVoucherReq) (model.Voucher, error)
	DeleteExpiredVoucher() error
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

func NewVoucherService(repo repository.VoucherRepo) VoucherService {
	return &voucherService{
		repo: repo,
	}
}
