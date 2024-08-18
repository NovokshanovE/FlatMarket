package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NovokshanovE/FlatMarket/internal/auth"
	"github.com/NovokshanovE/FlatMarket/internal/models"
	"github.com/NovokshanovE/FlatMarket/internal/services"
)

func CreateFlat(flatService *services.FlatService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var flat models.Flat
		if err := json.NewDecoder(r.Body).Decode(&flat); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		flat.Status = "created"

		if err := flatService.CreateFlat(&flat); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(flat)
	}
}

func UpdateFlat(flatService *services.FlatService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !auth.IsModerator(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var flat models.Flat
		if err := json.NewDecoder(r.Body).Decode(&flat); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := flatService.UpdateFlat(&flat); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(flat)
	}
}
