package texts

import (
	"backend/database"
	"backend/models"
)

func List(ClientID int64, LanguageID int64) ([]models.Text, error) {
	var texts []models.Text

	// TODO: filter by client
	if err := database.Database.
		Where("language_id = ?", LanguageID).
		Find(texts).Error; err != nil {
		return nil, err
	}

	return texts, nil
}