package domain

import "github.com/google/uuid"

// Клиент магазина
type Client struct {
	ID              uuid.UUID
	Name            string
	Surname         string
	Birthday        string
	Gender          string
	RegistrationDate string
	AddressID       uuid.UUID
}