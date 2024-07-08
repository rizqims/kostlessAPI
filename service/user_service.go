package service

import (
	"kostless-api/model"
	"kostless-api/repository"
	"kostless-api/util"
)

//interface
type UserServ interface{
	CreatedNewUser(payload model.User) (model.User, error)
}

//struct
type userServ struct {
	repo repository.UserRepo

}

//register implement
func (u *userServ) 	CreatedNewUser(payload model.User) (model.User, error) {
	hash, error := util.HashPassword(payload.Password)
	if error != nil {
		return model.User{}, error
	}
	payload.Password = hash
	return u.repo.CreatedNewUser(payload)
}

//constractor
func NewUserServ(reposi repository.UserRepo ) UserServ {
	return &userServ{repo: reposi}
}
