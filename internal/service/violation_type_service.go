package service

import (
	"errors"
	"violation-type-service/internal/model"
	"violation-type-service/internal/repository"
)

type violationTypeService struct {
	repo repository.ViolationTypeRepository
}

func NewViolationTypeService(r repository.ViolationTypeRepository) ViolationTypeService {
	return &violationTypeService{repo: r}
}

func (s *violationTypeService) GetAll() ([]model.ViolationType, error) {
	return s.repo.FindAll()
}

func (s *violationTypeService) GetByID(id int64) (model.ViolationType, error) {
	return s.repo.FindByID(id)
}

func (s *violationTypeService) Create(v model.ViolationType) (int64, error) {
	if v.Name == "" {
		return 0, errors.New("name is required")
	}
	if v.Name == "Другое" && v.OtherInfo == "" {
		return 0, errors.New("otherInfo is required for 'Другое'")
	}
	return s.repo.Create(v)
}

func (s *violationTypeService) Update(id int64, v model.ViolationType) error {
	if v.Name == "" {
		return errors.New("name is required")
	}
	return s.repo.Update(id, v)
}

func (s *violationTypeService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *violationTypeService) BulkInsert(list []model.ViolationType) error {
	return s.repo.BulkInsert(list)
}
