package service

import (
	"mysumselapi/internal/model"
	"mysumselapi/internal/repository"
)

type UtilityService interface {
	GetUtilities(userLat, userLon float64) (*model.UtilityResponse, error)
}

type utilityService struct {
	repo repository.UtilityRepository
}

func NewUtilityService(repo repository.UtilityRepository) UtilityService {
	return &utilityService{repo: repo}
}

func (s *utilityService) GetUtilities(userLat, userLon float64) (*model.UtilityResponse, error) {
	return s.repo.GetUtilities(userLat, userLon)
}
