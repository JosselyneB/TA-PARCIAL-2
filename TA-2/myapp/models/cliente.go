package services

import (
    "myapp/models"
    "myapp/repository"
)


type ClienteService struct {
    ClienteRepo *repository.ClienteRepository
}

func NewClienteService(clienteRepo *repository.ClienteRepository) *ClienteService {
    return &ClienteService{ClienteRepo: clienteRepo}
}

func (s *ClienteService) CreateCliente(cliente *models.Cliente) error {
    return s.ClienteRepo.CreateCliente(cliente)
}

func (s *ClienteService) GetClienteByID(id uint) (*models.Cliente, error) {
    return s.ClienteRepo.GetClienteByID(id)
}

func (s *ClienteService) UpdateCliente(cliente *models.Cliente) error {
    return s.ClienteRepo.UpdateCliente(cliente)
}

func (s *ClienteService) DeleteCliente(id uint) error {
    return s.ClienteRepo.DeleteCliente(id)
}

func (s *ClienteService) GetAllClientes() ([]models.Cliente, error) {
    return s.ClienteRepo.GetAllClientes()
}
