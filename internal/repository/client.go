package repository

import (
	"shopapi/internal/model"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ClientRepository interface {
	GetAll() []model.Client
}

type clientRepository struct {
	db *sqlx.DB
}

func NewClientRepository(db *sqlx.DB) ClientRepository {
	return &clientRepository{db: db}
}

// Пока заглушка, позже заменим на SELECT * FROM client
func (r *clientRepository) GetAll() []model.Client {
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
