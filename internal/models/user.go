package models

import "time"

var ModeratorsMap = map[string]int{}

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	UserType     string    `json:"user_type"`
	CreatedAt    time.Time `json:"created_at"`
}
