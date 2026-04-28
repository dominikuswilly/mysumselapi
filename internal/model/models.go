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
