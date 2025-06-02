package usecase

import (
	"shopapi/internal/domain"
	"shopapi/internal/repository"
)

type ClientUseCase struct {
	repo repository.ClientRepository
}

func NewClientUseCase(r repository.ClientRepository) *ClientUseCase {
	return &ClientUseCase{repo: r}
}

// Пример метода — пока пустой, логика будет позже.
func (uc *ClientUseCase) CreateClient(c domain.Client) error {
	// TODO: валидация и вызов repo.Create
	return nil
}
