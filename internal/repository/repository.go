package repository

import "violation-type-service/internal/model"

type ViolationRepository interface {
	FindAll() ([]model.ViolationType, error)
	Create(v model.ViolationType) (model.ViolationType, error)
	Update(id uint, v model.ViolationType) (model.ViolationType, error)
	Delete(id uint) error
	BulkInsert([]model.ViolationType) error
}