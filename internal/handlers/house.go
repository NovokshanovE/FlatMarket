package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NovokshanovE/FlatMarket/internal/auth"
	"github.com/NovokshanovE/FlatMarket/internal/models"
	"github.com/NovokshanovE/FlatMarket/internal/services"

	"github.com/gorilla/mux"
)

func CreateHouse(houseService *services.HouseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !auth.IsModerator(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var house models.House
		if err := json.NewDecoder(r.Body).Decode(&house); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := houseService.CreateHouse(&house); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(house)
	}
}

func GetFlatsByHouseID(houseService *services.HouseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		houseIDStr := vars["id"]
		houseID, err := strconv.Atoi(houseIDStr)
		if err != nil {
			http.Error(w, "Invalid house ID", http.StatusBadRequest)
			return
		}

		userType := auth.GetUserType(r)

		flats, err := houseService.GetFlatsByHouseID(houseID, userType)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"flats": flats,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
