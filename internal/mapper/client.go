package mapper

import (
	"shopapi/internal/domain"
	"shopapi/internal/dto"
	"time"

	"github.com/google/uuid"
)

func ToDomainClient(dto dto.CreateClientRequest) domain.Client {
	return domain.Client{
		ID:               uuid.New(),
		Name:             dto.Name,
		Surname:          dto.Surname,
		Birthday:         dto.Birthday,
		Gender:           dto.Gender,
		AddressID:        dto.AddressID,
		RegistrationDate: time.Now().Format("2006-01-02"),
	}
}

func ToClientResponse(c domain.Client) dto.ClientResponse {
	return dto.ClientResponse{
		ID:               c.ID,
		Name:             c.Name,
		Surname:          c.Surname,
		Birthday:         c.Birthday,
		Gender:           c.Gender,
		RegistrationDate: c.RegistrationDate,
		AddressID:        c.AddressID,
	}
}
