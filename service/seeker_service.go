package service

import (
	"fmt"
	"kostless-api/model"
	"kostless-api/model/dto"
	"kostless-api/repository"
	"kostless-api/util"
)

// interface
type SeekerServ interface {
	CreatedNewSeeker(payload model.Seekers) (model.Seekers, error)
	Login(payload dto.LoginDto) (dto.LoginResponse, error)
}

// struct
type seekerServ struct {
	repo repository.SeekerRepo
	jwt util.JwtToken
}

// Login implements SeekerServ.
func (s *seekerServ) Login(payload dto.LoginDto) (dto.LoginResponse, error) {
	seeker, err := s.repo.GetBySeeker(payload.Username)
	if err != nil {
		fmt.Print("err===",err)
		return dto.LoginResponse{}, fmt.Errorf("username invalid")
	}
	err = util.CheckPasswordHash(seeker.Password, payload.Password)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("password incorrect")
	}
	seeker.Password = ""
	token, err := s.jwt.GenerateToken(seeker.Username)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("password incorrect")
	}
	return token, nil
}

// register implement
func (s *seekerServ) CreatedNewSeeker(payload model.Seekers) (model.Seekers, error) {
	hash, error := util.HashPassword(payload.Password)
	if error != nil {
		return model.Seekers{}, error
	}
	payload.Password = hash
	return s.repo.CreatedNewSeeker(payload)
}

// constractor
func NewSeekerServ(reposi repository.SeekerRepo, jwt util.JwtToken) SeekerServ {
	return &seekerServ{repo: reposi, jwt: jwt}
}
