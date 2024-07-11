package mocks

import (
	"kostless/model"
	"kostless/model/dto"

	"github.com/stretchr/testify/mock"
)

type TransServiceMock struct {
	mock.Mock
}

func (t *TransServiceMock) CreateTrans(payload dto.TransCreateReq) (model.Trans, error) {
	args := t.Called(payload)
	return args.Get(0).(model.Trans), args.Error(1)
}

func (t *TransServiceMock) GetTransByID(id string) (model.Trans, error) {
	args := t.Called(id)
	return args.Get(0).(model.Trans), args.Error(1)
}

func (t *TransServiceMock) GetTransHistory() ([]model.Trans, error) {
	args := t.Called()
	return args.Get(0).([]model.Trans), args.Error(1)
}

func (t *TransServiceMock) GetPaylaterList() ([]model.Trans, error) {
	args := t.Called()
	return args.Get(0).([]model.Trans), args.Error(1)
}

func (t *TransServiceMock) GetTransByMonth(month, year string) ([]model.Trans, error) {
	args := t.Called(month, year)
	return args.Get(0).([]model.Trans), args.Error(1)
}

func (t *TransServiceMock) UpdatePaylater(payload dto.UpdatePaylaterReq) (dto.UpdatePaylaterRes, error) {
	args := t.Called(payload)
	return args.Get(0).(dto.UpdatePaylaterRes), args.Error(1)
}