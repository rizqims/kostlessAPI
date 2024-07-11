package mocks

import (
	"kostless/model"
	"kostless/model/dto"

	"github.com/stretchr/testify/mock"
)

type TransRepoMock struct {
	mock.Mock
}

func (t *TransRepoMock) CreateTrans(trans model.Trans) (model.Trans, error) {
	args := t.Called(trans)
	return args.Get(0).(model.Trans), args.Error(1)
}

func (t *TransRepoMock) GetTransByID(id string) (model.Trans, error) {
	args := t.Called(id)
	return args.Get(0).(model.Trans), args.Error(1)
}

func (t *TransRepoMock) GetTransHistory() ([]model.Trans, error) {
	args := t.Called()
	return args.Get(0).([]model.Trans), args.Error(1)
}

func (t *TransRepoMock) GetPaylaterList() ([]model.Trans, error) {
	args := t.Called()
	return args.Get(0).([]model.Trans), args.Error(1)
}

func (t *TransRepoMock) GetTransByMonth(startDate, endDate string) ([]model.Trans, error) {
	args := t.Called(startDate, endDate)
	return args.Get(0).([]model.Trans), args.Error(1)
}

func (t *TransRepoMock) UpdatePaylater(payload dto.UpdatePaylaterReq) (model.Trans, error) {
	args := t.Called(payload)
	return args.Get(0).(model.Trans), args.Error(1)
}