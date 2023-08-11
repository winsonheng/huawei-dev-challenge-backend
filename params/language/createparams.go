package language

import "backend/models"

type CreateParams struct {
	Name string `json:"name"`
}

func ToModel(params *CreateParams) *models.Language {
	return &models.Language{
		Name: params.Name,
	}
}