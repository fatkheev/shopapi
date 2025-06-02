package dto

import "github.com/google/uuid"

type CreateClientRequest struct {
	Name      string    `json:"name" binding:"required"`
	Surname   string    `json:"surname" binding:"required"`
	Birthday  string    `json:"birthday" binding:"required"`
	Gender    string    `json:"gender" binding:"required"`
	AddressID uuid.UUID `json:"address_id" binding:"required"`
}

type ClientResponse struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Surname          string    `json:"surname"`
	Birthday         string    `json:"birthday"`
	Gender           string    `json:"gender"`
	RegistrationDate string    `json:"registration_date"`
	AddressID        uuid.UUID `json:"address_id"`
}
