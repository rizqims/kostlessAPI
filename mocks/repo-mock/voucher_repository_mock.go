package mocks

import (
	"kostless/model"
	"kostless/model/dto"

	"github.com/stretchr/testify/mock"
)

type VoucherRepoMock struct {
	mock.Mock
}

func (v *VoucherRepoMock) CreateVoucher(payload dto.CreateVoucherReq) (model.Voucher, error) {
	args := v.Called(payload)
	return args.Get(0).(model.Voucher), args.Error(1)
}

func (v *VoucherRepoMock) DeleteExpiredVoucher() error {
	args := v.Called()
	return args.Error(0)
}

func (v *VoucherRepoMock) GetAllVoucher() ([]model.Voucher, error) {
	args := v.Called()
	return args.Get(0).([]model.Voucher), args.Error(1)
}

func (v *VoucherRepoMock) GetVoucherBySeekerID(id string) ([]model.Voucher, error) {
	args := v.Called(id)
	return args.Get(0).([]model.Voucher), args.Error(1)
}