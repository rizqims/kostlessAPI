package mocks

import (
	"kostless/model"
	"kostless/model/dto"

	"github.com/stretchr/testify/mock"
)

type SeekerServiceMock struct {
	mock.Mock
}

func (s *SeekerServiceMock) CreatedNewSeeker(payload model.Seekers) (model.Seekers, error) {
	args := s.Called(payload)
	return args.Get(0).(model.Seekers), args.Error(1)
}

func (s *SeekerServiceMock) Login(payload dto.LoginDto) (dto.LoginResponse, error) {
	args := s.Called(payload)
	return args.Get(0).(dto.LoginResponse), args.Error(1)
}

func (s *SeekerServiceMock) GetSeekerByID(id string) (model.Seekers, error) {
	args := s.Called(id)
	return args.Get(0).(model.Seekers), args.Error(1)
}

func (s *SeekerServiceMock) GetAllSeekers() ([]model.Seekers, error) {
	args := s.Called()
	return args.Get(0).([]model.Seekers), args.Error(1)
}

func (s *SeekerServiceMock) UpdateProfile(id string, updatedSeeker model.Seekers) error {
	args := s.Called(id, updatedSeeker)
	return args.Error(0)
}

func (s *SeekerServiceMock) DeleteSeeker(id string) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *SeekerServiceMock) UpdateAttitudePoints(id string, attitudePoints int) error {
	args := s.Called(id, attitudePoints)
	return args.Error(0)
}
