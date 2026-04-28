package repository

import (
	"mysumselapi/internal/model"
)

type CityRepository interface {
	GetAllCities(page, limit int) (*model.CityResponse, error)
}

type cityRepository struct{}

func NewCityRepository() CityRepository {
	return &cityRepository{}
}

func (r *cityRepository) GetAllCities(page, limit int) (*model.CityResponse, error) {
	strPtr := func(s string) *string { return &s }

	allCities := []model.City{
		{ID: 1, Name: "Kota Palembang", Type: "kota", Capital: nil, ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196536/ampera_uobaou.webp"},
		{ID: 2, Name: "Kota Lubuklinggau", Type: "kota", Capital: nil, ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/bukitserelo_hcsyjs.webp"},
		{ID: 3, Name: "Kota Pagar Alam", Type: "kota", Capital: nil, ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/gunungdempo_kg4mpt.webp"},
		{ID: 4, Name: "Kota Prabumulih", Type: "kota", Capital: nil, ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/megalith_ejctxk.webp"},
		{ID: 5, Name: "Kabupaten Ogan Komering Ulu", Type: "kabupaten", Capital: strPtr("Baturaja"), ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/goaputri_fgsbkb.webp"},
		{ID: 6, Name: "Kabupaten Ogan Komering Ulu Timur", Type: "kabupaten", Capital: strPtr("Martapura"), ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/curuptenang_movy9n.webp"},
		{ID: 7, Name: "Kabupaten Ogan Komering Ulu Selatan", Type: "kabupaten", Capital: strPtr("Muara Dua"), ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196536/danauranau_ior4ci.webp"},
		{ID: 8, Name: "Kabupaten Ogan Komering Ilir", Type: "kabupaten", Capital: strPtr("Kayu Agung"), ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196536/tamannasionalsembilang_r7s5dm.webp"},
		{ID: 9, Name: "Kabupaten Muara Enim", Type: "kabupaten", Capital: strPtr("Muara Enim"), ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/curuptenang_movy9n.webp"},
		{ID: 10, Name: "Kabupaten Lahat", Type: "kabupaten", Capital: strPtr("Lahat"), ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/bukitserelo_hcsyjs.webp"},
		{ID: 11, Name: "Kabupaten Musi Rawas", Type: "kabupaten", Capital: strPtr("Muara Beliti"), ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/megalith_ejctxk.webp"},
		{ID: 12, Name: "Kabupaten Musi Banyuasin", Type: "kabupaten", Capital: strPtr("Sekayu"), ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196536/tamannasionalsembilang_r7s5dm.webp"},
		{ID: 13, Name: "Kabupaten Banyuasin", Type: "kabupaten", Capital: strPtr("Pangkalan Balai"), ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196536/tamannasionalsembilang_r7s5dm.webp"},
		{ID: 14, Name: "Kabupaten Empat Lawang", Type: "kabupaten", Capital: strPtr("Tebing Tinggi"), ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/megalith_ejctxk.webp"},
		{ID: 15, Name: "Kabupaten Penukal Abab Lematang Ilir", Type: "kabupaten", Capital: strPtr("Talang Ubi"), ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/candibumiayu_qhalip.webp"},
		{ID: 16, Name: "Kabupaten Musi Rawas Utara", Type: "kabupaten", Capital: strPtr("Rupit"), ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/megalith_ejctxk.webp"},
		{ID: 17, Name: "Kabupaten Ogan Ilir", Type: "kabupaten", Capital: strPtr("Indralaya"), ImageURL: "https://res.cloudinary.com/dgixft7w5/image/upload/v1777196537/alquranakbar_ogwz8i.webp"},
	}

	total := len(allCities)
	offset := (page - 1) * limit
	if offset > total {
		offset = total
	}

	end := offset + limit
	if end > total {
		end = total
	}

	paginatedCities := allCities[offset:end]

	resp := &model.CityResponse{
		Data: paginatedCities,
	}
	resp.Meta.Total = total
	resp.Meta.Provinsi = "Sumatera Selatan"
	resp.Meta.Page = page
	resp.Meta.Limit = limit

	return resp, nil
}
