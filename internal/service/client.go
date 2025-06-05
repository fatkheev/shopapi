package service

import (
	"shopapi/internal/model"
	"shopapi/internal/repository"
)

type ClientService interface {
	GetAllClients() []model.Client
}

type clientService struct {
	repo repository.ClientRepository
}

func NewClientService(repo repository.ClientRepository) ClientService {
	return &clientService{repo: repo}
}

func (s *clientService) GetAllClients() []model.Client {
	return s.repo.GetAll()
}
