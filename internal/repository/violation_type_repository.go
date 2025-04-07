package repository

import (
	"gorm.io/gorm"
	"violation-type-service/internal/model"
)

type violationRepo struct {
	db *gorm.DB
}

func NewViolationRepository(db *gorm.DB) ViolationRepository {
	return &violationRepo{db: db}
}

func (r *violationRepo) FindAll() ([]model.ViolationType, error) {
	var list []model.ViolationType
	res := r.db.Find(&list)
	return list, res.Error
}

func (r *violationRepo) Create(v model.ViolationType) (model.ViolationType, error) {
	res := r.db.Create(&v)
	return v, res.Error
}

func (r *violationRepo) Update(id uint, v model.ViolationType) (model.ViolationType, error) {
	var existing model.ViolationType
	if err := r.db.First(&existing, id).Error; err != nil {
		return model.ViolationType{}, err
	}
	existing.Name = v.Name
	existing.OtherInfo = v.OtherInfo
	if err := r.db.Save(&existing).Error; err != nil {
		return model.ViolationType{}, err
	}
	return existing, nil
}

func (r *violationRepo) Delete(id uint) error {
	return r.db.Delete(&model.ViolationType{}, id).Error
}

func (r *violationRepo) BulkInsert(list []model.ViolationType) error {
	return r.db.Create(&list).Error
}