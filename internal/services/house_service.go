package services

import (
	"database/sql"
	"time"

	"github.com/NovokshanovE/FlatMarket/internal/models"
)

type HouseService struct {
	DB *sql.DB
}

func (hs *HouseService) CreateHouse(house *models.House) error {
	query := `
        INSERT INTO houses (address, year, developer, created_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at
    `
	err := hs.DB.QueryRow(query, house.Address, house.Year, house.Developer, time.Now()).Scan(&house.ID, &house.CreatedAt)
	return err
}

func (hs *HouseService) GetFlatsByHouseID(houseID int, userType string) ([]models.Flat, error) {
	query := `
        SELECT id, house_id, price, rooms, status, created_at
        FROM flats
        WHERE house_id = $1
    `
	if userType == "client" {
		query += " AND status = 'approved'"
	}

	rows, err := hs.DB.Query(query, houseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flats []models.Flat
	for rows.Next() {
		var flat models.Flat
		if err := rows.Scan(&flat.ID, &flat.HouseID, &flat.Price, &flat.Rooms, &flat.Status, &flat.CreatedAt); err != nil {
			return nil, err
		}
		flats = append(flats, flat)
	}
	return flats, nil
}
