package service

import (
	"kostless/model"
	"kostless/model/dto"
	"kostless/repository"
)

type KosService interface {
	CreateKos(request dto.KosRequest) (model.Kos, error)
	UpdateKos(id string, request dto.KosRequest) (model.Kos, error)
	DeleteKos(id string) error
	GetKosByID(id string) (model.Kos, error)
}

type kosService struct {
	repo repository.KosRepository
}

func NewKosService(repo repository.KosRepository) *kosService {
	return &kosService{repo}
}

func (s *kosService) CreateKos(request dto.KosRequest) (model.Kos, error) {
	kos := model.Kos{
		Name:        request.Name,
		Address:     request.Address,
		RoomCount:   request.RoomCount,
		Coordinate:  request.Coordinate,
		Description: request.Description,
		Rules:       request.Rules,
	}

	return s.repo.CreateKos(kos)
}

func (s *kosService) UpdateKos(id string, request dto.KosRequest) (model.Kos, error) {
	kos := model.Kos{
		ID:          id,
		Name:        request.Name,
		Address:     request.Address,
		RoomCount:   request.RoomCount,
		Coordinate:  request.Coordinate,
		Description: request.Description,
		Rules:       request.Rules,
	}

	return s.repo.UpdateKos(kos)
}

func (s *kosService) DeleteKos(id string) error {
	return s.repo.DeleteKos(id)
}

func (s *kosService) GetKosByID(id string) (model.Kos, error) {
	return s.repo.GetKosByID(id)
}
