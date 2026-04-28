package repository

import (
	"mysumselapi/internal/model"
)

type HeroRepository interface {
	GetHeroes() ([]model.Hero, error)
}

type heroRepository struct{}

func NewHeroRepository() HeroRepository {
	return &heroRepository{}
}

func (r *heroRepository) GetHeroes() ([]model.Hero, error) {
	return []model.Hero{
		{ID: "1", ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777345926/hero1_c4nr55.jpg"},
		{ID: "2", ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777345926/hero2_shejrj.jpg"},
		{ID: "3", ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777345931/hero3_jize6z.jpg"},
		{ID: "4", ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777345926/hero4_nlxz93.jpg"},
		{ID: "5", ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777345926/hero5_mnpf7v.jpg"},
	}, nil
}
