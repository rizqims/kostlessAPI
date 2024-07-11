package mocks

import (
	"kostless/model"

	"github.com/stretchr/testify/mock"
)

type RoomRepoMock struct {
	mock.Mock
}

func (r *RoomRepoMock) CreateRoom(room model.Room) (model.Room, error) {
	args := r.Called(room)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomRepoMock) GetAllRooms() ([]model.Room, error) {
	args := r.Called()
	return args.Get(0).([]model.Room), args.Error(1)
}

func (r *RoomRepoMock) GetRoomByID(id string) (model.Room, error) {
	args := r.Called(id)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomRepoMock) GetRoomByAvailability(availability string) ([]model.Room, error) {
	args := r.Called(availability)
	return args.Get(0).([]model.Room), args.Error(1)
}

func (r *RoomRepoMock) GetRoomByPriceLowerThanOrEqual(price int) ([]model.Room, error) {
	args := r.Called(price)
	return args.Get(0).([]model.Room), args.Error(1)
}

func (r *RoomRepoMock) UpdateRoom(room model.Room) (model.Room, error) {
	args := r.Called(room)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomRepoMock) DeleteRoom(id string) error {
	args := r.Called(id)
	return args.Error(0)
}