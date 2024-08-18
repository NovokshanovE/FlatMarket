package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NovokshanovE/FlatMarket/internal/auth"
)

func DummyLogin(w http.ResponseWriter, r *http.Request) {
	userType := r.URL.Query().Get("user_type")
	if userType != "client" && userType != "moderator" {
		http.Error(w, "Invalid user type", http.StatusBadRequest)
		return
	}

	token := auth.GenerateToken(userType)
	response := map[string]string{"token": token}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
