package services

import (
	"database/sql"
	"time"

	"github.com/NovokshanovE/FlatMarket/internal/models"
)

// FlatService provides operations on flats.
type FlatService struct {
	DB *sql.DB
}

// CreateFlat creates a new flat in the database.
func (fs *FlatService) CreateFlat(flat *models.Flat) error {
	query := `
        INSERT INTO flats (house_id, price, rooms, status, created_at)
        VALUES ($1, $2, $3, 'created', $4)
        RETURNING id, created_at
    `
	err := fs.DB.QueryRow(query, flat.HouseID, flat.Price, flat.Rooms, time.Now()).Scan(&flat.ID, &flat.CreatedAt)
	if err != nil {
		return err
	}

	// Update the last_flat_added column in the houses table
	updateQuery := `
        UPDATE houses SET last_flat_added = $1 WHERE id = $2
    `
	_, err = fs.DB.Exec(updateQuery, time.Now(), flat.HouseID)
	return err
}

// UpdateFlat updates an existing flat in the database.
func (fs *FlatService) UpdateFlat(flat *models.Flat) error {
	query := `
        UPDATE flats SET price = $1, rooms = $2, status = $3 WHERE id = $4
        RETURNING created_at
    `
	err := fs.DB.QueryRow(query, flat.Price, flat.Rooms, flat.Status, flat.ID).Scan(&flat.CreatedAt)
	return err
}
