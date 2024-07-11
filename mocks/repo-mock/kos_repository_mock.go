package repomock
import (
	"kostless/model"

	"github.com/stretchr/testify/mock"
)

type KosRepoMock struct {
	mock.Mock
}

func (k *KosRepoMock) CreateKos(kos model.Kos) (model.Kos, error) {
	args := k.Called(kos)
	return args.Get(0).(model.Kos), args.Error(1)
}

func (k *KosRepoMock) UpdateKos(kos model.Kos) (model.Kos, error) {
	args := k.Called(kos)
	return args.Get(0).(model.Kos), args.Error(1)
}

func (k *KosRepoMock) DeleteKos(id string) error {
	args := k.Called(id)
	return args.Error(0)
}

func (k *KosRepoMock) GetKosByID(id string) (model.Kos, error) {
	args := k.Called(id)
	return args.Get(0).(model.Kos), args.Error(1)
}
