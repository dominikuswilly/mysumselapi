package repository

import (
	"math"
	"mysumselapi/internal/model"
	"time"
)

type UtilityRepository interface {
	GetUtilities(userLat, userLon float64) (*model.UtilityResponse, error)
}

type utilityRepository struct{}

func NewUtilityRepository() UtilityRepository {
	return &utilityRepository{}
}

func (r *utilityRepository) GetUtilities(userLat, userLon float64) (*model.UtilityResponse, error) {
	items := []model.UtilityItem{
		{
			ID:        "hosp-plm-001",
			Category:  "hospital",
			Name:      "Siloam Hospitals Sriwijaya",
			City:      "Palembang",
			IsOpen24h: true,
			ImageURL:  "https://res.cloudinary.com/dgixft7w5/image/upload/v1777358342/siloamsriwijaya_frhxkw.jpg",
			Contact: model.UtilityContact{
				Phone:     "+627115229100",
				Emergency: "1500911",
			},
			Location: model.UtilityLocation{
				Address:   "Jl. POM IX, Lorok Pakjo, Ilir Barat I",
				Latitude:  -2.977696,
				Longitude: 104.742350,
			},
			Tags: []string{"International Staff", "ICU", "Specialist"},
		},
		{
			ID:        "hosp-llg-002",
			Category:  "hospital",
			Name:      "RSUD Siti Aisyah Lubuk Linggau",
			City:      "Lubuk Linggau",
			IsOpen24h: true,
			ImageURL:  "https://res.cloudinary.com/dgixft7w5/image/upload/v1777358342/sitiaisyah_hbwbvf.webp",
			Contact: model.UtilityContact{
				Phone:     "+62733451604",
				Emergency: "+628117190022",
			},
			Location: model.UtilityLocation{
				Address:   "Jl. Lapter Silampari No.20, Air Kuti",
				Latitude:  -3.3101,
				Longitude: 102.8821,
			},
			Tags: []string{"Public Hospital", "Emergency Unit"},
		},
		{
			ID:        "pol-plm-001",
			Category:  "police",
			Name:      "Polrestabes Palembang",
			City:      "Palembang",
			IsOpen24h: true,
			ImageURL:  "https://res.cloudinary.com/dgixft7w5/image/upload/v1777358342/polrestabespalembang_ovv6zk.webp",
			Contact: model.UtilityContact{
				Phone:     "+62711510455",
				Emergency: "110",
			},
			Location: model.UtilityLocation{
				Address:   "Jl. Gub H Bastari No. 1, Jakabaring",
				Latitude:  -3.0039,
				Longitude: 104.7684,
			},
			Tags: []string{"Tourism Police Desk", "24h Support"},
		},
		{
			ID:        "mc-plm-001",
			Category:  "money_changer",
			Name:      "Remaja Money Changer (Authorized)",
			City:      "Palembang",
			IsOpen24h: false,
			ImageURL:  "https://res.cloudinary.com/dgixft7w5/image/upload/v1777358342/remajamoney_xfk16u.jpg",
			Contact: model.UtilityContact{
				Phone: "+62711311145",
			},
			Location: model.UtilityLocation{
				Address:   "Jl. Jend. Sudirman No. 98/456 (Simpang Charitas)",
				Latitude:  -2.9745,
				Longitude: 104.7558,
			},
			Metadata: &model.UtilityMetadata{
				OperatingHours: "08:00 - 17:00",
				AuthorizedID:   "PVA/12/SS/2024",
			},
		},
	}

	// Calculate distances if coordinates are provided
	for i := range items {
		if userLat != 0 || userLon != 0 {
			items[i].Location.DistanceKM = calculateDistance(userLat, userLon, items[i].Location.Latitude, items[i].Location.Longitude)
		} else {
			// Mock distances if no coordinates provided (optional, or just leave as 0)
			items[i].Location.DistanceKM = 0.0
		}
	}

	resp := &model.UtilityResponse{
		Items: items,
	}
	resp.Meta.TotalResults = len(items)
	resp.Meta.NearestCity = "Palembang"
	resp.Meta.ServerTime = time.Now()

	return resp, nil
}

// calculateDistance using Haversine formula
func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Earth radius in km
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return math.Round(R*c*10) / 10 // Round to 1 decimal place
}
