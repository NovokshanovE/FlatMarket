package models

import "time"

type Flat struct {
	ID        int       `json:"id"`
	HouseID   int       `json:"house_id"`
	Price     int       `json:"price"`
	Rooms     int       `json:"rooms"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
