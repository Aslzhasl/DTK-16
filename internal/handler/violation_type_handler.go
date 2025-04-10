package handler

import (
	"net/http"
	"strconv"

	"violation-type-service/internal/excel"
	"violation-type-service/internal/model"
	"violation-type-service/internal/repository"
	"violation-type-service/internal/service"

	"github.com/gin-gonic/gin"
)

type ViolationTypeHandler struct {
	service service.ViolationTypeService
	repo    repository.ViolationTypeRepository
}

func NewViolationTypeHandler(s service.ViolationTypeService, r repository.ViolationTypeRepository) *ViolationTypeHandler {
	return &ViolationTypeHandler{service: s, repo: r}
}

func (h *ViolationTypeHandler) GetAll(c *gin.Context) {
	data, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *ViolationTypeHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	data, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *ViolationTypeHandler) Create(c *gin.Context) {
	var req model.ViolationType
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *ViolationTypeHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var req model.ViolationType
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.Update(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (h *ViolationTypeHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
func (h *ViolationTypeHandler) ImportExcel(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		path = "./data.xlsx"
	}
	err := excel.ImportFromExcel(path, h.repo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Импорт завершен"})
}
