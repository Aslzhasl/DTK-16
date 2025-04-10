package service

import "violation-type-service/internal/model"

type ViolationTypeService interface {
	GetAll() ([]model.ViolationType, error)
	GetByID(id int64) (model.ViolationType, error)
	Create(v model.ViolationType) (int64, error)
	Update(id int64, v model.ViolationType) error
	Delete(id int64) error
	BulkInsert([]model.ViolationType) error
}
