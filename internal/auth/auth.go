package auth

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func GenerateToken(userType string) string {
	return uuid.New().String() + "_" + userType
}

func IsModerator(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return false
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return false
	}

	token := parts[1]
	return strings.HasSuffix(token, "_moderator")
}

func GetUserType(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	token := parts[1]
	if strings.HasSuffix(token, "_moderator") {
		return "moderator"
	}
	if strings.HasSuffix(token, "_client") {
		return "client"
	}
	return ""
}
