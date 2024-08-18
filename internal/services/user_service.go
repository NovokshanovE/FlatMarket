package services

import (
	"database/sql"
	"errors"
	"time"

	"github.com/NovokshanovE/FlatMarket/internal/auth"
	"github.com/NovokshanovE/FlatMarket/internal/models"
)

// UserService provides operations on users.
type UserService struct {
	DB *sql.DB
}

// RegisterUser registers a new user in the database.
func (us *UserService) RegisterUser(user *models.User) error {
	query := `
        INSERT INTO users (email, password_hash, user_type, created_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at
    `
	err := us.DB.QueryRow(query, user.Email, user.PasswordHash, user.UserType, time.Now()).Scan(&user.ID, &user.CreatedAt)
	return err
}

// LoginUser authenticates a user and returns a token.
func (us *UserService) LoginUser(email, password string) (string, error) {
	var user models.User
	query := `
        SELECT id, password_hash, user_type FROM users WHERE email = $1
    `
	err := us.DB.QueryRow(query, email).Scan(&user.ID, &user.PasswordHash, &user.UserType)
	if err != nil {
		return "", err
	}

	// Check password (using a simple comparison for demonstration; use a secure hash comparison in production)
	if user.PasswordHash != password {
		return "", errors.New("invalid credentials")
	}

	// Generate token (use a proper token generation mechanism in production)
	token := auth.GenerateToken(user.UserType)
	return token, nil
}
