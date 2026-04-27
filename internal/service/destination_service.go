package service

import (
	"mysumselapi/internal/model"
	"mysumselapi/internal/repository"
)

type DestinationService interface {
	GetFavoriteDestinations() ([]model.Destination, error)
}

type destinationService struct {
	repo repository.DestinationRepository
}

func NewDestinationService(repo repository.DestinationRepository) DestinationService {
	return &destinationService{repo: repo}
}

func (s *destinationService) GetFavoriteDestinations() ([]model.Destination, error) {
	return s.repo.GetFavoriteDestinations()
}
