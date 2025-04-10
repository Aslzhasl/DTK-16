package middleware

import (
	"violation-type-service/internal/auth"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(authClient auth.AuthClient, requiredRole string) gin.HandlerFunc {
	return auth.JWTMiddleware(authClient, requiredRole)
}
