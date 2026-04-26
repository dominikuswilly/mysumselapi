package service

import (
	"mysumselapi/internal/model"
	"mysumselapi/internal/repository"
)

type CityService interface {
	GetAllCities() (*model.CityResponse, error)
}

type cityService struct {
	repo repository.CityRepository
}

func NewCityService(repo repository.CityRepository) CityService {
	return &cityService{repo: repo}
}

func (s *cityService) GetAllCities() (*model.CityResponse, error) {
	return s.repo.GetAllCities()
}
