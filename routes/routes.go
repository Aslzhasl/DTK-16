package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"violation-type-service/internal/handler"
	"violation-type-service/internal/repository"
	"violation-type-service/internal/service"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	repo := repository.NewViolationRepository(db)
	service := service.NewViolationService(repo)
	handler := handler.NewViolationHandler(service, repo)

	r.GET("/api/violations", handler.GetAll)
	r.POST("/api/violations", handler.Create)
	r.PUT("/api/violations/:id", handler.Update)
	r.DELETE("/api/violations/:id", handler.Delete)
	r.POST("/api/violations/import", handler.ImportExcel)

	return r
}