package middleware

import (
	"context"
	"net/http"
	"strings"

	"ecom-local/internal/auth"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "userID"
const UserEmailKey contextKey = "userEmail"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Jeton d'authentification manquant", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Format du header Authorization invalide. Format attendu: Bearer <token>", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		token, err := auth.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Jeton invalide ou expiré", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Claims invalides dans le jeton", http.StatusUnauthorized)
			return
		}

		// Injecter l'utilisateur dans le Context
		ctx := r.Context()
		if sub, ok := claims["sub"].(float64); ok {
			ctx = context.WithValue(ctx, UserIDKey, int64(sub))
		}
		if email, ok := claims["email"].(string); ok {
			ctx = context.WithValue(ctx, UserEmailKey, email)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
