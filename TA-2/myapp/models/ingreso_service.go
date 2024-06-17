package services

import (
    "myapp/models"
    "myapp/repository"

)

type IngresoService struct {
    IngresoRepo *repository.IngresoRepository
}

func NewIngresoService(ingresoRepo *repository.IngresoRepository) *IngresoService {
    return &IngresoService{IngresoRepo: ingresoRepo}
}

func (s *IngresoService) CreateIngreso(ingreso *models.Ingreso) error {
    return s.IngresoRepo.CreateIngreso(ingreso)
}

func (s *IngresoService) GetIngresoByID(id uint) (*models.Ingreso, error) {
    return s.IngresoRepo.GetIngresoByID(id)
}

func (s *IngresoService) UpdateIngreso(ingreso *models.Ingreso) error {
    return s.IngresoRepo.UpdateIngreso(ingreso)
}

func (s *IngresoService) DeleteIngreso(id uint) error {
    return s.IngresoRepo.DeleteIngreso(id)
}

func (s *IngresoService) GetAllIngresos() ([]models.Ingreso, error) {
    return s.IngresoRepo.GetAllIngresos()
}
