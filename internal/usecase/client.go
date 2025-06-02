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

func (uc *ClientUseCase) CreateClient(c domain.Client) error {
	// Пока заглушка — будет логика позже
	return nil
}
