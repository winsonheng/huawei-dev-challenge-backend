package texts

import (
	"backend/database"
	"backend/models"

	"gorm.io/gorm"
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

func Create(Text *models.Text, ClientText *models.ClientText) (*models.Text, error) {
	err := database.Database.Transaction(func(db *gorm.DB) error {
		// check whether text exists
		// create if not
		res := db.Model(&models.Text{}).
			Where("language_id = ?", Text.LanguageID).
			Where("content = ?", Text.Content).
			Find(&Text)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			if err := db.Create(&Text).Error; err != nil {
				return err
			}
		}

		ClientText.TextID = Text.ID

		return db.Create(&ClientText).Error
	})

	if err != nil {
		return nil, err
	}

	return Text, nil
}