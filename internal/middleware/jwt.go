package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	ClaimsKey contextKey = "jwt_claims"
)

func JWTAdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Получение секрета
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			http.Error(w, "JWT secret not configured", http.StatusInternalServerError)
			return
		}

		// 2. Проверка заголовка Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// 3. Парсинг токена
		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		if tokenString == "" {
			http.Error(w, "Bearer token required", http.StatusUnauthorized)
			return
		}

		// 4. Валидация токена
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		// 5. Обработка ошибок парсинга
		if err != nil {
			http.Error(w, fmt.Sprintf("Token validation failed: %v", err), http.StatusUnauthorized)
			return
		}

		// 6. Проверка валидности токена
		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// 7. Извлечение claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// 8. Проверка роли
		role, ok := claims["role"].(string)
		if !ok || role != "ROLE_ADMIN" {
			http.Error(w, "Insufficient privileges", http.StatusForbidden)
			return
		}

		// 9. Добавление claims в контекст
		ctx := context.WithValue(r.Context(), ClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
