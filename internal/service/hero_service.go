package service

import (
	"mysumselapi/internal/model"
	"mysumselapi/internal/repository"
)

type HeroService interface {
	GetHeroes() ([]model.Hero, error)
}

type heroService struct {
	repo repository.HeroRepository
}

func NewHeroService(repo repository.HeroRepository) HeroService {
	return &heroService{repo: repo}
}

func (s *heroService) GetHeroes() ([]model.Hero, error) {
	return s.repo.GetHeroes()
}
