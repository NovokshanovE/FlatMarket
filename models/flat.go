package models

import (
	"time"

	"gorm.io/gorm"
)

type Flat struct {
	gorm.Model
	HouseID   uint      `json:"house_id"`
	Number    int       `json:"number"`
	Price     int       `json:"price"`
	Rooms     int       `json:"rooms"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
