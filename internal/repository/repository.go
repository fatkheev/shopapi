package repository

import "github.com/google/uuid"
import "shopapi/internal/domain"

// Интерфейс для работы с клиентами.
type ClientRepository interface {
	Create(client domain.Client) error
	Delete(id uuid.UUID) error
	GetByID(id uuid.UUID) (domain.Client, error)
	GetAll(limit, offset int) ([]domain.Client, error)
	UpdateAddress(id uuid.UUID, newAddressID uuid.UUID) error
}
