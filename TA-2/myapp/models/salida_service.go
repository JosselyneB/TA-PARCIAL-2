package services

import (
     "myapp/models"
    "myapp/repository"
)

type SalidaService struct {
    SalidaRepo *repository.SalidaRepository
}

func NewSalidaService(salidaRepo *repository.SalidaRepository) *SalidaService {
    return &SalidaService{SalidaRepo: salidaRepo}
}

func (s *SalidaService) CreateSalida(salida *models.Salida) error {
    return s.SalidaRepo.CreateSalida(salida)
}

func (s *SalidaService) GetSalidaByID(id uint) (*models.Salida, error) {
    return s.SalidaRepo.GetSalidaByID(id)
}

func (s *SalidaService) UpdateSalida(salida *models.Salida) error {
    return s.SalidaRepo.UpdateSalida(salida)
}

func (s *SalidaService) DeleteSalida(id uint) error {
    return s.SalidaRepo.DeleteSalida(id)
}

func (s *SalidaService) GetAllSalidas() ([]models.Salida, error) {
    return s.SalidaRepo.GetAllSalidas()
}
