package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClient interface {
	VerifyUser(token string) (UserInfo, error)
}

type UserInfo struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
	Email   string `json:"email"`
	Role    string `json:"role"`
}

func JWTWithAuth(authClient AuthClient, requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 1. Получение токена
			tokenString, err := extractToken(r)
			if err != nil {
				respondError(w, http.StatusUnauthorized, err.Error())
				return
			}

			// 2. Валидация JWT
			_, err = validateJWT(tokenString)
			if err != nil {
				respondError(w, http.StatusUnauthorized, "Invalid token")
				return
			}

			// 3. Проверка пользователя в сервисе аутентификации
			userInfo, err := authClient.VerifyUser(tokenString)
			if err != nil {
				respondError(w, http.StatusUnauthorized, "User verification failed")
				return
			}

			// 4. Проверка роли
			if userInfo.Role != "ROLE_ADMIN" {
				respondError(w, http.StatusForbidden, "Insufficient privileges")
				return
			}

			// 5. Проверка, что токен действителен
			if !userInfo.Valid {
				respondError(w, http.StatusUnauthorized, "Token is not valid: "+userInfo.Message)
				return
			}

			// 6. Добавление информации в контекст
			ctx := context.WithValue(r.Context(), "user", userInfo)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func validateJWT(tokenString string) (jwt.MapClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization header required")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", fmt.Errorf("invalid authorization header format")
	}

	return parts[1], nil
}

func respondError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
