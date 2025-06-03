package model

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID               uuid.UUID
	Name             string
	Surname          string
	Birthday         time.Time
	Gender           string
	RegistrationDate time.Time
	AddressID        uuid.UUID
}

type ClientResponseDTO struct {
	ID               string `json:"id"`
	Name             string `json:"name,omitempty"`
	Surname          string `json:"surname,omitempty"`
	Birthday         string `json:"birthday,omitempty"`
	Gender           string `json:"gender,omitempty"`
	RegistrationDate string `json:"registration_date,omitempty"`
}
