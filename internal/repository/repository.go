package repository

import "violation-type-service/internal/model"

type ViolationTypeRepository interface {
	FindAll() ([]model.ViolationType, error)
	FindByID(id int64) (model.ViolationType, error)
	Create(v model.ViolationType) (int64, error)
	Update(id int64, v model.ViolationType) error
	Delete(id int64) error
	BulkInsert([]model.ViolationType) error
}
