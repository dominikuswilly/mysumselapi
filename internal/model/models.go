package model

import "time"

// Response represents a standard API response structure
type Response struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    any       `json:"data,omitempty"`
	Time    time.Time `json:"time"`
}

// Message represents a simple domain entity
type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

// City represents a city or regency in South Sumatra
type City struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Capital  *string `json:"capital"`
	ImageURL string  `json:"image_url"`
}

// CityResponse represents the specific response structure for cities
type CityResponse struct {
	Meta struct {
		Total    int    `json:"total"`
		Provinsi string `json:"provinsi"`
		Page     int    `json:"page"`
		Limit    int    `json:"limit"`
	} `json:"meta"`
	Data []City `json:"data"`
}

// Category represents a destination category
type Category struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// Destination represents a tourist destination
type Destination struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Location   string     `json:"location"`
	Categories []Category `json:"categories"`
	ImageURL   string     `json:"image_url"`
}

// Hero represents a hero banner item
type Hero struct {
	ID       string `json:"id"`
	ImageURL string `json:"image_url"`
}

// Utility models
type UtilityContact struct {
	Phone     string `json:"phone,omitempty"`
	Emergency string `json:"emergency,omitempty"`
}

type UtilityLocation struct {
	Address    string  `json:"address"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	DistanceKM float64 `json:"distance_km"`
}

type UtilityMetadata struct {
	OperatingHours string `json:"operating_hours,omitempty"`
	AuthorizedID   string `json:"authorized_id,omitempty"`
}

type UtilityItem struct {
	ID        string           `json:"id"`
	Category  string           `json:"category"`
	Name      string           `json:"name"`
	City      string           `json:"city"`
	IsOpen24h bool             `json:"is_open_24h"`
	ImageURL  string           `json:"image_url"`
	Contact   UtilityContact   `json:"contact"`
	Location  UtilityLocation  `json:"location"`
	Tags      []string         `json:"tags,omitempty"`
	Metadata  *UtilityMetadata `json:"metadata,omitempty"`
}

type UtilityResponse struct {
	Items []UtilityItem `json:"items"`
	Meta  struct {
		TotalResults int       `json:"total_results"`
		NearestCity  string    `json:"nearest_city"`
		ServerTime   time.Time `json:"server_time"`
	} `json:"meta"`
}
