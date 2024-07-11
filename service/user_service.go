package service

import (
	"fmt"
	"kostless/model"
	"kostless/model/dto"
	"kostless/repository"
	"kostless/util"
	"time"

	"github.com/google/uuid"

)

// interface
type UserServ interface {
	CreatedNewUser(payload model.User) (model.User, error)
	Login(payload dto.LoginDto) (dto.LoginResponse, error)
	UpdateProfile(id string, updatedUser model.User) error
	GetUser(id string) (model.User, error)
}

// struct
type userServ struct {
	repo repository.UserRepo
	jwt  util.JwtToken
}

// GetUser implements UserServ.
func (u *userServ) GetUser(id string) (model.User, error) {
	return u.repo.GetUserById(id)
}

// UpdateProfile implements UserServ.
func (u *userServ) UpdateProfile(id string, updatedUser model.User) error {
	updatedUser.Id = id
	updatedUser.UpdatedAt = time.Now()
	return u.repo.PutUpdateUserProf(id, updatedUser)
}

// Login implements UserServ.
func (u *userServ) Login(payload dto.LoginDto) (dto.LoginResponse, error) {
	user, err := u.repo.GetByUsername(payload.Username)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("username invalid")
	}
	err = util.CheckPasswordHash(user.Password, payload.Password)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("password incorrect")
	}
	token, err := u.jwt.GenerateToken(user.Id, user.Username)
	if err != nil {
		fmt.Print("errr ===", err)
		return dto.LoginResponse{}, fmt.Errorf("failed generate token")

	}
	return token, nil
}

// register implement
func (u *userServ) CreatedNewUser(payload model.User) (model.User, error) {
	payload.Id = uuid.New().String()
	payload.UpdatedAt = time.Now()
	hash, error := util.HashPassword(payload.Password)
	if error != nil {
		return model.User{}, error
	}
	payload.Password = hash
	return u.repo.CreatedNewUser(payload)
}

// constractor
func NewUserServ(reposi repository.UserRepo, jwt util.JwtToken) UserServ {
	return &userServ{repo: reposi, jwt: jwt}
}
