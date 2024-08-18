package auth

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	tests := []struct {
		userType string
		want     string
	}{
		{"client", "token_client"},
		{"moderator", "token_moderator"},
	}

	for _, tt := range tests {
		t.Run(tt.userType, func(t *testing.T) {
			if got := GenerateToken(tt.userType); got != tt.want {
				t.Errorf("GenerateToken(%v) = %v, want %v", tt.userType, got, tt.want)
			}
		})
	}
}
