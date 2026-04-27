package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"mysumselapi/internal/model"
	"mysumselapi/internal/repository"
	"mysumselapi/internal/service"
)

func main() {
	// Initialize Repositories
	msgRepo := repository.NewMessageRepository()
	cityRepo := repository.NewCityRepository()
	destRepo := repository.NewDestinationRepository()

	// Initialize Services
	msgService := service.NewMessageService(msgRepo)
	cityService := service.NewCityService(cityRepo)
	destService := service.NewDestinationService(destRepo)

	// Initialize Router
	mux := http.NewServeMux()

	// Health Check
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		sendJSON(w, http.StatusOK, model.Response{
			Status:  "success",
			Message: "Service is healthy",
			Time:    time.Now(),
		})
	})

	// Welcome Endpoint
	mux.HandleFunc("GET /api/v1/welcome", func(w http.ResponseWriter, r *http.Request) {
		msg, err := msgService.GetWelcome()
		if err != nil {
			sendJSON(w, http.StatusInternalServerError, model.Response{
				Status:  "error",
				Message: err.Error(),
				Time:    time.Now(),
			})
			return
		}
		sendJSON(w, http.StatusOK, model.Response{
			Status:  "success",
			Message: "Welcome message retrieved",
			Data:    msg,
			Time:    time.Now(),
		})
	})

	// Cities Endpoint
	mux.HandleFunc("GET /api/v1/cities", func(w http.ResponseWriter, r *http.Request) {
		cities, err := cityService.GetAllCities()
		if err != nil {
			sendJSON(w, http.StatusInternalServerError, model.Response{
				Status:  "error",
				Message: err.Error(),
				Time:    time.Now(),
			})
			return
		}
		sendJSON(w, http.StatusOK, model.Response{
			Status:  "success",
			Message: "Cities retrieved successfully",
			Data:    cities,
			Time:    time.Now(),
		})
	})

	// Favorite Destinations Endpoint
	mux.HandleFunc("GET /api/v1/destinations/favorites", func(w http.ResponseWriter, r *http.Request) {
		favorites, err := destService.GetFavoriteDestinations()
		if err != nil {
			sendJSON(w, http.StatusInternalServerError, model.Response{
				Status:  "error",
				Message: err.Error(),
				Time:    time.Now(),
			})
			return
		}
		sendJSON(w, http.StatusOK, model.Response{
			Status:  "success",
			Message: "Favorite destinations retrieved successfully",
			Data:    favorites,
			Time:    time.Now(),
		})
	})

	// Start Server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      loggerMiddleware(mux),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Server starting on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func sendJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}
