package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/sakhamoori/mytolet/api/internal/auth"
)

type contextKey string

const UserContextKey contextKey = "user"

func AuthMiddleware(authProvider *auth.JWTProvider) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the Authorization header
			authHeader := r.Header.Get("Authorization")
			
			// Skip auth for certain paths
			if strings.HasPrefix(r.URL.Path, "/health") || 
			   (strings.HasPrefix(r.URL.Path, "/query") && r.Method == "GET") {
				next.ServeHTTP(w, r)
				return
			}
			
			// Check if header is present and has correct format
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				// Proceed without authentication (some operations don't require it)
				next.ServeHTTP(w, r)
				return
			}
			
			// Extract the token
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			
			// Validate the token
			claims, err := authProvider.ValidateToken(tokenString)
			if err != nil {
				// Token is invalid, but we'll still let GraphQL handle auth
				next.ServeHTTP(w, r)
				return
			}
			
			// Set user in context
			ctx := context.WithValue(r.Context(), UserContextKey, claims)
			
			// Call the next handler with the updated context
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}