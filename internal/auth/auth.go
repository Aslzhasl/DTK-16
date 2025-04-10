package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
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

func ValidateJWT(tokenString string) error {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return fmt.Errorf("JWT_SECRET is not set")
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return fmt.Errorf("invalid token: %v", err)
	}
	return nil
}

func JWTMiddleware(authClient AuthClient, requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")

		if err := ValidateJWT(token); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		userInfo, err := authClient.VerifyUser(token)
		if err != nil || !userInfo.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User verification failed"})
			return
		}
		if userInfo.Role != requiredRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		ctx := context.WithValue(c.Request.Context(), "user", userInfo)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
