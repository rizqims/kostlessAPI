package service

import (
	"fmt"
	"kostless-api/model"
	"kostless-api/model/dto"
	"kostless-api/repository"
	"kostless-api/util"
)

// interface
type UserServ interface {
	CreatedNewUser(payload model.User) (model.User, error)
	Login(payload dto.LoginDto) (dto.LoginResponse, error)
}

// struct
type userServ struct {
	repo repository.UserRepo
	jwt util.JwtToken
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
	token, err := u.jwt.GenerateToken(user.Username)
	if err != nil {
		fmt.Print("errr ===",err)
		return dto.LoginResponse{}, fmt.Errorf("password incorrect")
	}
	return token, nil
}

// register implement
func (u *userServ) CreatedNewUser(payload model.User) (model.User, error) {
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
