package service

import (
	"shopapi/internal/model"
	"time"

	"github.com/google/uuid"
)

type ClientService interface {
	GetAllClients() []model.Client
}

type clientService struct {}

func NewClientService() ClientService {
	return &clientService{}
}

func (s *clientService) GetAllClients() []model.Client {
	client := model.Client{
		ID:               uuid.New(),
		Name:             "Иван",
		Surname:          "Иванов",
		Birthday:         time.Date(1990, 5, 12, 0, 0, 0, 0, time.UTC),
		Gender:           "male",
		RegistrationDate: time.Now(),
		AddressID:        uuid.New(),
	}

	return []model.Client{client}
}
