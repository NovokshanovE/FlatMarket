package models

import (
	"time"

	"gorm.io/gorm"
)

type House struct {
	gorm.Model
	Address         string    `json:"address"`
	YearBuilt       int       `json:"year_built"`
	Developer       string    `json:"developer"`
	CreatedAt       time.Time `json:"created_at"`
	LastFlatAddedAt time.Time `json:"last_flat_added_at"`
}
