package mocks

import (
	"kostless/model"
	"kostless/model/dto"

	"github.com/stretchr/testify/mock"
)

type KostServiceMock struct {
	mock.Mock
}

func (k *KostServiceMock) CreateKos(request dto.KosRequest) (model.Kos, error) {
	args := k.Called(request)
	return args.Get(0).(model.Kos), args.Error(1)
}

func (k *KostServiceMock) UpdateKos(id string, request dto.KosRequest) (model.Kos, error) {
	args := k.Called(id, request)
	return args.Get(0).(model.Kos), args.Error(1)
}

func (k *KostServiceMock) DeleteKos(id string) error {
	args := k.Called(id)
	return args.Error(0)
}

func (k *KostServiceMock) GetKosByID(id string) (model.Kos, error) {
	args := k.Called(id)
	return args.Get(0).(model.Kos), args.Error(1)
}