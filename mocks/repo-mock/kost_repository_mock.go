package mocks

import (
	"kostless/model"

	"github.com/stretchr/testify/mock"
)

type KostRepoMock struct {
	mock.Mock
}

func (k *KostRepoMock) CreateKos(kos model.Kos) (model.Kos, error) {
	args := k.Called(kos)
	return args.Get(0).(model.Kos), args.Error(1)
}

func (k *KostRepoMock) UpdateKos(kos model.Kos) (model.Kos, error) {
	args := k.Called(kos)
	return args.Get(0).(model.Kos), args.Error(1)
}

func (k *KostRepoMock) DeleteKos(id string) error {
	args := k.Called(id)
	return args.Error(0)
}

func (k *KostRepoMock) GetKosByID(id string) (model.Kos, error) {
	args := k.Called(id)
	return args.Get(0).(model.Kos), args.Error(1)
}
