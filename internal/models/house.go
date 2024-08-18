package models

import "time"

type House struct {
	ID            int        `json:"id"`
	Address       string     `json:"address"`
	Year          int        `json:"year"`
	Developer     *string    `json:"developer,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	LastFlatAdded *time.Time `json:"last_flat_added,omitempty"`
}
