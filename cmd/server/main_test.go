package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/NovokshanovE/FlatMarket/internal/handlers"
	"github.com/NovokshanovE/FlatMarket/internal/services"

	"github.com/gorilla/mux"
)

func setupRouter(houseService *services.HouseService, flatService *services.FlatService) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/dummyLogin", handlers.DummyLogin).Methods("GET")
	r.HandleFunc("/house/create", handlers.CreateHouse(houseService)).Methods("POST")
	r.HandleFunc("/house/{id:[0-9]+}", handlers.GetFlatsByHouseID(houseService)).Methods("GET")
	r.HandleFunc("/flat/create", handlers.CreateFlat(flatService)).Methods("POST")
	r.HandleFunc("/flat/update", handlers.UpdateFlat(flatService)).Methods("POST")
	return r
}

func TestDummyLogin(t *testing.T) {
	req, err := http.NewRequest("GET", "/dummyLogin?user_type=client", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.DummyLogin)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"token":"token_client"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
