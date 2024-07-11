package service

import (
	"kostless/model"
	"kostless/model/dto"
	"kostless/repository"
	"strconv"
)

type RoomService interface {
	CreateRoom(request dto.RoomRequest) (model.Room, error)
	GetAllRooms() ([]model.Room, error)
	GetRoomByID(id string) (model.Room, error)
	GetRoomByAvailability(availability string) ([]model.Room, error)
	GetRoomByPriceLowerThanOrEqual(budget string) ([]model.Room, error)
	UpdateRoom(room model.Room) (model.Room, error)
	DeleteRoom(id string) error
}

type roomService struct {
	roomRepository repository.RoomRepository
}

func NewRoomService(roomRepository repository.RoomRepository) *roomService {
	return &roomService{roomRepository}
}

func (s *roomService) CreateRoom(request dto.RoomRequest) (model.Room, error) {
	room := model.Room{
		KosID:       request.KosID,
		Name:        request.Name,
		Type:        request.Type,
		Description: request.Description,
		Avail:       request.Avail,
		Price:       request.Price,
	}
	return s.roomRepository.CreateRoom(room)
}

func (s *roomService) UpdateRoom(room model.Room) (model.Room, error) {
	return s.roomRepository.UpdateRoom(room)
}

func (s *roomService) GetAllRooms() ([]model.Room, error) {
	return s.roomRepository.GetAllRooms()
}

func (s *roomService) GetRoomByID(id string) (model.Room, error) {
	return s.roomRepository.GetRoomByID(id)
}

func (s *roomService) GetRoomByAvailability(availability string) ([]model.Room, error) {
	return s.roomRepository.GetRoomByAvailability(availability)
}

func (s *roomService) GetRoomByPriceLowerThanOrEqual(budget string) ([]model.Room, error) {
	price, err := strconv.Atoi(budget)
	if err != nil {
		return nil, err
	}
	return s.roomRepository.GetRoomByPriceLowerThanOrEqual(price)
}

func (s *roomService) DeleteRoom(id string) error {
	return s.roomRepository.DeleteRoom(id)
}
