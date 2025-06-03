package mapper

import "shopapi/internal/model"

func MapClientToResponseDTO(c model.Client) model.ClientResponseDTO {
	return model.ClientResponseDTO{
		ID:               c.ID.String(),
		Name:             c.Name,
		Surname:          c.Surname,
		Birthday:         c.Birthday.Format("2006-01-02"),
		Gender:           c.Gender,
		RegistrationDate: c.RegistrationDate.Format("2006-01-02"),
	}
}
