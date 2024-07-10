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
type SeekerServ interface {
	CreatedNewSeeker(payload model.Seekers) (model.Seekers, error)
	Login(payload dto.LoginDto) (dto.LoginResponse, error)
	GetSeekerByID(id string) (model.Seekers, error)
	GetAllSeekers() ([]model.Seekers, error)
	UpdateProfile(id string, updatedSeeker model.Seekers) error
	DeleteSeeker(id string) error
	UpdateAttitudePoints(id string, attitudePoints int) error
}

// struct
type seekerServ struct {
	repo repository.SeekerRepo
	jwt  util.JwtToken
}

// DeleteSeeker implements SeekerServ.
func (s *seekerServ) DeleteSeeker(id string) error {
	return s.repo.DeleteSeeker(id)
}

// GetAllSeekers implements SeekerServ.
func (s *seekerServ) GetAllSeekers() ([]model.Seekers, error) {
	return s.repo.GetAllSeekers()
}

// GetSeekerByID implements SeekerServ.
func (s *seekerServ) GetSeekerByID(id string) (model.Seekers, error) {
	return s.repo.GetSeekerByID(id)
}

// UpdateAttitudePoints implements SeekerServ.
func (s *seekerServ) UpdateAttitudePoints(id string, attitudePoints int) error {
	seeker, err := s.repo.GetSeekerByID(id)
	if err != nil {
		return err
	}
	if attitudePoints < 5 {
		if err := util.SendEmail(seeker.Email, "Low Attitude Points", "Your attitude points are below 5."); err != nil {
			return err
		}
	} else if attitudePoints > 10 {
		if err := util.NotifyOwner("Add vouchers to seeker with ID " + id); err != nil {
			return err
		}
	}
	return s.repo.UpdateAttitudePoints(id, attitudePoints)
}

// UpdateProfile implements SeekerServ.
func (s *seekerServ) UpdateProfile(id string, updatedSeeker model.Seekers) error {
	updatedSeeker.Id = id
	updatedSeeker.UpdatedAt = time.Now()
	return s.repo.UpdateSeeker(id, updatedSeeker)
}

// Login implements SeekerServ.
func (s *seekerServ) Login(payload dto.LoginDto) (dto.LoginResponse, error) {
	seeker, err := s.repo.GetBySeeker(payload.Username)
	if err != nil {
		fmt.Print("err===", err)
		return dto.LoginResponse{}, fmt.Errorf("username invalid")
	}
	err = util.CheckPasswordHash(seeker.Password, payload.Password)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("password incorrect")
	}
	seeker.Password = ""
	token, err := s.jwt.GenerateToken(seeker.Id, seeker.Username)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("failed generate token")
	}
	return token, nil
}

// register implement
func (s *seekerServ) CreatedNewSeeker(payload model.Seekers) (model.Seekers, error) {
	payload.Id = uuid.New().String()
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
