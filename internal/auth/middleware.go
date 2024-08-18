package auth

import (
	"net/http"
	"strings"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token := parts[1]
		userType := parseToken(token)
		if userType == "" {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// You can set userType in the request context if needed
		// ctx := context.WithValue(r.Context(), "userType", userType)
		// next.ServeHTTP(w, r.WithContext(ctx))

		next.ServeHTTP(w, r)
	})
}

func parseToken(token string) string {
	parts := strings.Split(token, "_")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}
