package service

import (
	"kostless-api/model"
	"kostless-api/repository"
	"kostless-api/util"
)

//interface
type SeekerServ interface{
	CreatedNewSeeker(payload model.Seekers) (model.Seekers, error)
}

//struct
type seekerServ struct {
	repo repository.SeekerRepo

}

//register implement
func (s *seekerServ) 	CreatedNewSeeker(payload model.Seekers) (model.Seekers, error) {
	hash, error := util.HashPassword(payload.Password)
	if error != nil {
		return model.Seekers{}, error
	}
	payload.Password = hash
	return s.repo.CreatedNewSeeker(payload)
}

//constractor
func NewSeekerServ(reposi repository.SeekerRepo ) SeekerServ {
	return &seekerServ{repo: reposi}
}
