package servicemock

import (
	"kostless/model"
	"kostless/model/dto"

	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (u *UserServiceMock) CreatedNewUser(payload model.User) (model.User, error) {
	args := u.Called(payload)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserServiceMock) Login(payload dto.LoginDto) (dto.LoginResponse, error) {
	args := u.Called(payload)
	return args.Get(0).(dto.LoginResponse), args.Error(1)
}

func (u *UserServiceMock) UpdateProfile(id string, updatedUser model.User) error {
	args := u.Called(id, updatedUser)
	return args.Error(0)
}

func (u *UserServiceMock) GetUser(id string) (model.User, error) {
	args := u.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}
