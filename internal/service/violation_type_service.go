package service

import (
	"errors"
	"violation-type-service/internal/model"
	"violation-type-service/internal/repository"
)

type violationService struct {
	repo repository.ViolationRepository
}

func NewViolationService(r repository.ViolationRepository) ViolationService {
	return &violationService{repo: r}
}

func (s *violationService) GetAll() ([]model.ViolationType, error) {
	return s.repo.FindAll()
}

func (s *violationService) Create(v model.ViolationType) (model.ViolationType, error) {
	if v.Name == "Другое" && v.OtherInfo == "" {
		return model.ViolationType{}, errors.New("Если тип 'Другое' — нужно указать дополнительную информацию")
	}
	return s.repo.Create(v)
}

func (s *violationService) Update(id uint, v model.ViolationType) (model.ViolationType, error) {
	return s.repo.Update(id, v)
}

func (s *violationService) Delete(id uint) error {
	return s.repo.Delete(id)
}