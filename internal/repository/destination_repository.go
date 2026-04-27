package repository

import (
	"mysumselapi/internal/model"
)

type DestinationRepository interface {
	GetFavoriteDestinations() ([]model.Destination, error)
}

type destinationRepository struct{}

func NewDestinationRepository() DestinationRepository {
	return &destinationRepository{}
}

func (r *destinationRepository) GetFavoriteDestinations() ([]model.Destination, error) {
	return []model.Destination{
		{
			ID:       "dest-001",
			Name:     "Taman Nasional Sembilang",
			Location: "Banyuasin",
			Categories: []model.Category{
				{ID: "cat_1", Slug: "nature", Name: "Nature"},
				{ID: "cat_5", Slug: "adventure", Name: "Adventure"},
			},
			ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196536/tamannasionalsembilang_r7s5dm.webp",
		},
		{
			ID:       "dest-002",
			Name:     "Gua Putri",
			Location: "Ogan Komering Ulu",
			Categories: []model.Category{
				{ID: "cat_1", Slug: "nature", Name: "Nature"},
				{ID: "cat_2", Slug: "culture", Name: "Culture & Heritage"},
			},
			ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/goaputri_fgsbkb.webp",
		},
		{
			ID:       "dest-003",
			Name:     "Curup Tenang",
			Location: "Muara Enim",
			Categories: []model.Category{
				{ID: "cat_1", Slug: "nature", Name: "Nature"},
				{ID: "cat_7", Slug: "waterfall", Name: "Waterfall"},
			},
			ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/curuptenang_movy9n.webp",
		},
		{
			ID:       "dest-004",
			Name:     "Al-Quran Al-Akbar",
			Location: "Palembang",
			Categories: []model.Category{
				{ID: "cat_2", Slug: "culture", Name: "Culture"},
				{ID: "cat_6", Slug: "religious", Name: "Religious"},
			},
			ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/alquranakbar_ogwz8i.webp",
		},
		{
			ID:       "dest-005",
			Name:     "Gunung Dempo",
			Location: "Pagar Alam",
			Categories: []model.Category{
				{ID: "cat_1", Slug: "nature", Name: "Nature"},
				{ID: "cat_8", Slug: "hiking", Name: "Hiking"},
			},
			ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/gunungdempo_kg4mpt.webp",
		},
		{
			ID:       "dest-006",
			Name:     "Situs Megalith",
			Location: "Lahat",
			Categories: []model.Category{
				{ID: "cat_2", Slug: "culture", Name: "Culture & Heritage"},
				{ID: "cat_9", Slug: "history", Name: "History"},
			},
			ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/megalith_ejctxk.webp",
		},
		{
			ID:       "dest-007",
			Name:     "Candi Bumi Ayu",
			Location: "PALI",
			Categories: []model.Category{
				{ID: "cat_2", Slug: "culture", Name: "Culture & Heritage"},
				{ID: "cat_9", Slug: "history", Name: "History"},
			},
			ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/candibumiayu_qhalip.webp",
		},
		{
			ID:       "dest-008",
			Name:     "Bukit Serelo",
			Location: "Lahat",
			Categories: []model.Category{
				{ID: "cat_1", Slug: "nature", Name: "Nature"},
				{ID: "cat_10", Slug: "landmark", Name: "Landmark"},
			},
			ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/bukitserelo_hcsyjs.webp",
		},
		{
			ID:       "dest-009",
			Name:     "Jembatan Ampera",
			Location: "Palembang",
			Categories: []model.Category{
				{ID: "cat_10", Slug: "landmark", Name: "Landmark"},
				{ID: "cat_3", Slug: "culinary", Name: "Culinary"},
			},
			ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196536/ampera_uobaou.webp",
		},
		{
			ID:       "dest-010",
			Name:     "Danau Ranau",
			Location: "OKU Selatan",
			Categories: []model.Category{
				{ID: "cat_1", Slug: "nature", Name: "Nature"},
				{ID: "cat_4", Slug: "recreation", Name: "Recreation"},
			},
			ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196536/danauranau_ior4ci.webp",
		},
	}, nil
}
