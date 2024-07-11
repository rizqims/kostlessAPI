package mocks

import (
	"kostless/model"
	"kostless/model/dto"

	"github.com/stretchr/testify/mock"
)

type VoucherServiceMock struct {
	mock.Mock
}

func (v *VoucherServiceMock) CreateVoucher(payload dto.CreateVoucherReq) (model.Voucher, error) {
	args := v.Called(payload)
	return args.Get(0).(model.Voucher), args.Error(1)
}

func (v *VoucherServiceMock) DeleteExpiredVoucher() error {
	args := v.Called()
	return args.Error(0)
}

func (v *VoucherServiceMock) GetAllVoucher() ([]model.Voucher, error) {
	args := v.Called()
	return args.Get(0).([]model.Voucher), args.Error(1)
}

func (v *VoucherServiceMock) GetVoucherBySeekerID(id string) ([]model.Voucher, error) {
	args := v.Called(id)
	return args.Get(0).([]model.Voucher), args.Error(1)
}
