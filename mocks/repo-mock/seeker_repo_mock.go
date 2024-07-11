package mocks

import (
	"kostless/model"

	"github.com/stretchr/testify/mock"
)

type SeekerRepoMock struct {
	mock.Mock
}

func (s *SeekerRepoMock) DeleteSeeker(id string) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *SeekerRepoMock) GetAllSeekers() ([]model.Seekers, error) {
	args := s.Called()
	return args.Get(0).([]model.Seekers), args.Error(1)
}

func (s *SeekerRepoMock) GetSeekerByID(id string) (model.Seekers, error) {
	args := s.Called(id)
	return args.Get(0).(model.Seekers), args.Error(1)
}

func (s *SeekerRepoMock) UpdateAttitudePoints(id string, attitudePoints int) error {
	args := s.Called(id, attitudePoints)
	return args.Error(0)
}

func (s *SeekerRepoMock) UpdateSeeker(id string, seeker model.Seekers) error {
	args := s.Called(id, seeker)
	return args.Error(0)
}

func (s *SeekerRepoMock) GetBySeeker(username string) (model.Seekers, error) {
	args := s.Called(username)
	return args.Get(0).(model.Seekers), args.Error(1)
}

func (s *SeekerRepoMock) GetByID(id string) (model.Seekers, error) {
	args := s.Called(id)
	return args.Get(0).(model.Seekers), args.Error(1)
}