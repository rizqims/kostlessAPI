package service

import (
	"kostless/model"
	"kostless/model/dto"
	"kostless/repository"
)

type KosService interface {
	CreateKos(request dto.KosRequest) (model.Kos, error)
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
