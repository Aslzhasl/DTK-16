package service

import "violation-type-service/internal/model"

type ViolationService interface {
	GetAll() ([]model.ViolationType, error)
	Create(v model.ViolationType) (model.ViolationType, error)
	Update(id uint, v model.ViolationType) (model.ViolationType, error)
	Delete(id uint) error
}