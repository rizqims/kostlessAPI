package repomock

import (
	"kostless/model"

	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (u *UserRepoMock) GetUserById(id string) (model.User, error) {
	args := u.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserRepoMock) PutUpdateUserProf(id string, user model.User) error {
	args := u.Called(id, user)
	return args.Error(0)
}

func (u *UserRepoMock) GetByUsername(username string) (model.User, error) {
	args := u.Called(username)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserRepoMock) CreatedNewUser(payload model.User) (model.User, error) {
	args := u.Called(payload)
	return args.Get(0).(model.User), args.Error(1)
}
