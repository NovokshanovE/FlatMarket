package auth

import (
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestGenerateToken(t *testing.T) {
	userTypes := []string{"client", "moderator"}

	for _, userType := range userTypes {
		t.Run(userType, func(t *testing.T) {
			token := GenerateToken(userType)
			if !strings.HasSuffix(token, "_"+userType) {
				t.Errorf("Expected token to end with _%s, got %s", userType, token)
			}

			// Check if the generated part before "_" is a valid UUID
			uuidPart := strings.Split(token, "_")[0]
			if _, err := uuid.Parse(uuidPart); err != nil {
				t.Errorf("Expected token to contain valid UUID, got error: %v", err)
			}
		})
	}
}

func TestIsModerator(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)

	moderatorToken := GenerateToken("moderator")
	req.Header.Set("Authorization", "Bearer "+moderatorToken)

	if !IsModerator(req) {
		t.Error("Expected IsModerator to return true")
	}

	clientToken := GenerateToken("client")
	req.Header.Set("Authorization", "Bearer "+clientToken)
	if IsModerator(req) {
		t.Error("Expected IsModerator to return false")
	}
}
