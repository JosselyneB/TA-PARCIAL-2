package services

import (
    "myapp/models"
    "myapp/repository"
)

type PagoService struct {
    PagoRepo *repository.PagoRepository
}

func NewPagoService(pagoRepo *repository.PagoRepository) *PagoService {
    return &PagoService{PagoRepo: pagoRepo}
}

func (s *PagoService) CreatePago(pago *models.Pago) error {
    return s.PagoRepo.CreatePago(pago)
}

func (s *PagoService) GetPagoByID(id uint) (*models.Pago, error) {
    return s.PagoRepo.GetPagoByID(id)
}

func (s *PagoService) UpdatePago(pago *models.Pago) error {
    return s.PagoRepo.UpdatePago(pago)
}

func (s *PagoService) DeletePago(id uint) error {
    return s.PagoRepo.DeletePago(id)
}

func (s *PagoService) GetAllPagos() ([]models.Pago, error) {
    return s.PagoRepo.GetAllPagos()
}
