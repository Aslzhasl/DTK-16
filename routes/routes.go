package routes

import (
	"violation-type-service/internal/auth"
	"violation-type-service/internal/handler"
	"violation-type-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.ViolationTypeHandler) *gin.Engine {
	r := gin.Default()

	authClient := auth.NewJavaAuthClient("http://localhost:8081")
	adminRoutes := r.Group("/api/violation-types")
	adminRoutes.Use(middleware.JWTMiddleware(authClient, "ROLE_ADMIN"))
	{
		adminRoutes.GET("", h.GetAll)
		adminRoutes.GET(":id", h.GetByID)
		adminRoutes.POST("", h.Create)
		adminRoutes.PUT(":id", h.Update)
		adminRoutes.DELETE(":id", h.Delete)
		adminRoutes.POST("/import", h.ImportExcel)
	}

	return r
}
