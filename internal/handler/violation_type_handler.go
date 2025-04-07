package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"violation-type-service/internal/model"
	"violation-type-service/internal/service"
	"violation-type-service/internal/excel"
	"violation-type-service/internal/repository"
	"os"
)

type ViolationHandler struct {
	service service.ViolationService
	repo    repository.ViolationRepository
}

func NewViolationHandler(s service.ViolationService, r repository.ViolationRepository) *ViolationHandler {
	return &ViolationHandler{service: s, repo: r}
}

func (h *ViolationHandler) GetAll(c *gin.Context) {
	list, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *ViolationHandler) Create(c *gin.Context) {
	var input model.ViolationType
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.service.Create(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *ViolationHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input model.ViolationType
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.service.Update(uint(id), input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *ViolationHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *ViolationHandler) ImportExcel(c *gin.Context) {
	filePath := c.Query("path")
	if filePath == "" {
		filePath = "./violations.xlsx"
	}
	if _, err := os.Stat(filePath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Файл не найден"})
		return
	}
	err := excel.ImportFromExcel(filePath, h.repo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Импорт завершен"})
}
