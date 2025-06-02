package repository

import (
	"shopapi/internal/domain"
	"github.com/google/uuid"
)

type ClientRepository interface {
	Create(client domain.Client) error
	Delete(id uuid.UUID) error
	GetByID(id uuid.UUID) (domain.Client, error)
	GetAll(limit, offset int) ([]domain.Client, error)
	UpdateAddress(id uuid.UUID, newAddressID uuid.UUID) error
}
