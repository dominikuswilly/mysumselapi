package service

import (
	"mysumselapi/internal/model"
	"mysumselapi/internal/repository"
)

type CityService interface {
	GetAllCities(page, limit int) (*model.CityResponse, error)
}

type cityService struct {
	repo repository.CityRepository
}

func NewCityService(repo repository.CityRepository) CityService {
	return &cityService{repo: repo}
}

func (s *cityService) GetAllCities(page, limit int) (*model.CityResponse, error) {
	return s.repo.GetAllCities(page, limit)
}
