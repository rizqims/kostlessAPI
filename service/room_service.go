package service

import (
	"kostless/model"
	"kostless/model/dto"
	"kostless/repository"
	"strconv"

	"github.com/google/uuid"
)

type RoomService interface {
	CreateRoom(request dto.RoomRequest) (model.Room, error)
	GetRoomByID(id string) (model.Room, error)
	GetRoomByAvailability(availability string) ([]model.Room, error)
	GetRoomByPriceLowerThan(budget string) ([]model.Room, error)
}

type roomService struct {
	roomRepository repository.RoomRepository
}

func NewRoomService(roomRepository repository.RoomRepository) *roomService {
	return &roomService{roomRepository}
}

func (s *roomService) CreateRoom(request dto.RoomRequest) (model.Room, error) {
	room := model.Room{
		Name:        request.Name,
		Type:        request.Type,
		Description: request.Description,
		Avail:       request.Avail,
		Price:       request.Price,
	}
	return s.roomRepository.CreateRoom(room)
}

func (s *roomService) GetRoomByID(id string) (model.Room, error) {
	newId, err := uuid.Parse(id)
	if err != nil {
		return model.Room{}, err
	}
	return s.roomRepository.GetRoomByID(newId)
}

func (s *roomService) GetRoomByAvailability(availability string) ([]model.Room, error) {
	return s.roomRepository.GetRoomByAvailability(availability)
}

func (s *roomService) GetRoomByPriceLowerThan(budget string) ([]model.Room, error) {
	price, err := strconv.Atoi(budget)
	if err != nil {
		return nil, err
	}
	return s.roomRepository.GetRoomByPriceLowerThan(price)
}
