package main

import (
	"log"
	"net/http"
	"os"

	"github.com/NovokshanovE/FlatMarket/internal/auth"
	"github.com/NovokshanovE/FlatMarket/internal/config"
	"github.com/NovokshanovE/FlatMarket/internal/database"
	"github.com/NovokshanovE/FlatMarket/internal/handlers"
	"github.com/NovokshanovE/FlatMarket/internal/services"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	db, err := database.Connect(&cfg.Database)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	houseService := &services.HouseService{DB: db}
	flatService := &services.FlatService{DB: db}

	r := mux.NewRouter()

	r.Handle("/dummyLogin", http.HandlerFunc(handlers.DummyLogin)).Methods("GET")
	r.Handle("/house/create", auth.AuthorizationMiddleware(http.HandlerFunc(handlers.CreateHouse(houseService)))).Methods("POST")
	r.Handle("/house/{id:[0-9]+}", auth.AuthorizationMiddleware(http.HandlerFunc(handlers.GetFlatsByHouseID(houseService)))).Methods("GET")
	r.Handle("/flat/create", auth.AuthorizationMiddleware(http.HandlerFunc(handlers.CreateFlat(flatService)))).Methods("POST")
	r.Handle("/flat/update", auth.AuthorizationMiddleware(http.HandlerFunc(handlers.UpdateFlat(flatService)))).Methods("POST")

	handler := cors.Default().Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Server.Port
	}
	log.Printf("Server is starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
