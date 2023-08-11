package languages

import (
	"backend/database"
	"backend/models"
)

func Create(language *models.Language) (*models.Language, error) {
	if err := database.Database.Create(&language).Error; err != nil {
		return nil, err
	}
	return language, nil
}

func List() ([]models.Language, error) {
	var langauges []models.Language
	if err := database.Database.Find(&langauges).Error; err != nil {
		return nil, err
	}
	return langauges, nil
}