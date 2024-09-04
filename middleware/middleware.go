package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// jwtKEY is the secret key used for signing JWT tokens
var jwtKey = []byte("your_secret_key")

// contextKey is a type used to avoid context key collisions
type contextKey string

// UserContextKey is the key used to store user info in the context
const UserContextKey contextKey = "user"

// Claims represents the structure of the JWT claims, it includes the user ID and standard JWT claims
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

// JWTAuthMiddleware is a middleware function that checks the validity of the JWT token and adds user information to the request context.
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header from the request
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Extract the token from the Authorization header
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}

		// Parse the JWT token and extract claims
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			// Ensure the token method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtKey, nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Check if the token is valid
		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add user information to context
		ctx := context.WithValue(r.Context(), UserContextKey, claims)
		// Pass the request with the updated context to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
