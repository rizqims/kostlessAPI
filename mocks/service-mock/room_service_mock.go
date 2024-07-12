package servicemock

import (
	"kostless/model"
	"kostless/model/dto"

	"github.com/stretchr/testify/mock"
)

type RoomServiceMock struct {
	mock.Mock
}

func (r *RoomServiceMock) CreateRoom(request dto.RoomRequest) (model.Room, error) {
	args := r.Called(request)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomServiceMock) GetAllRooms() ([]model.Room, error) {
	args := r.Called()
	return args.Get(0).([]model.Room), args.Error(1)
}

func (r *RoomServiceMock) GetRoomByID(id string) (model.Room, error) {
	args := r.Called(id)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomServiceMock) GetRoomByAvailability(availability string) ([]model.Room, error) {
	args := r.Called(availability)
	return args.Get(0).([]model.Room), args.Error(1)
}

func (r *RoomServiceMock) GetRoomByPriceLowerThanOrEqual(budget string) ([]model.Room, error) {
	args := r.Called(budget)
	return args.Get(0).([]model.Room), args.Error(1)
}

func (r *RoomServiceMock) UpdateRoom(room model.Room) (model.Room, error) {
	args := r.Called(room)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomServiceMock) DeleteRoom(id string) error {
	args := r.Called(id)
	return args.Error(0)
}
