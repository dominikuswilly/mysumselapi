package repository

import (
	"mysumselapi/internal/model"
)

type CityRepository interface {
	GetAllCities() (*model.CityResponse, error)
}

type cityRepository struct{}

func NewCityRepository() CityRepository {
	return &cityRepository{}
}

func (r *cityRepository) GetAllCities() (*model.CityResponse, error) {
	strPtr := func(s string) *string { return &s }

	cities := []model.City{
		{ID: 1, Name: "Kota Palembang", Type: "kota", Capital: nil},
		{ID: 2, Name: "Kota Lubuklinggau", Type: "kota", Capital: nil},
		{ID: 3, Name: "Kota Pagar Alam", Type: "kota", Capital: nil},
		{ID: 4, Name: "Kota Prabumulih", Type: "kota", Capital: nil},
		{ID: 5, Name: "Kabupaten Ogan Komering Ulu", Type: "kabupaten", Capital: strPtr("Baturaja")},
		{ID: 6, Name: "Kabupaten Ogan Komering Ulu Timur", Type: "kabupaten", Capital: strPtr("Martapura")},
		{ID: 7, Name: "Kabupaten Ogan Komering Ulu Selatan", Type: "kabupaten", Capital: strPtr("Muara Dua")},
		{ID: 8, Name: "Kabupaten Ogan Komering Ilir", Type: "kabupaten", Capital: strPtr("Kayu Agung")},
		{ID: 9, Name: "Kabupaten Muara Enim", Type: "kabupaten", Capital: strPtr("Muara Enim")},
		{ID: 10, Name: "Kabupaten Lahat", Type: "kabupaten", Capital: strPtr("Lahat")},
		{ID: 11, Name: "Kabupaten Musi Rawas", Type: "kabupaten", Capital: strPtr("Muara Beliti")},
		{ID: 12, Name: "Kabupaten Musi Banyuasin", Type: "kabupaten", Capital: strPtr("Sekayu")},
		{ID: 13, Name: "Kabupaten Banyuasin", Type: "kabupaten", Capital: strPtr("Pangkalan Balai")},
		{ID: 14, Name: "Kabupaten Empat Lawang", Type: "kabupaten", Capital: strPtr("Tebing Tinggi")},
		{ID: 15, Name: "Kabupaten Penukal Abab Lematang Ilir", Type: "kabupaten", Capital: strPtr("Talang Ubi")},
		{ID: 16, Name: "Kabupaten Musi Rawas Utara", Type: "kabupaten", Capital: strPtr("Rupit")},
		{ID: 17, Name: "Kabupaten Ogan Ilir", Type: "kabupaten", Capital: strPtr("Indralaya")},
	}

	resp := &model.CityResponse{
		Data: cities,
	}
	resp.Meta.Total = len(cities)
	resp.Meta.Provinsi = "Sumatera Selatan"

	return resp, nil
}
