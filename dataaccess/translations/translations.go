package translations

import (
	"backend/database"
	"backend/models"
	"errors"

	"gorm.io/gorm"
)

func Create(SourceText *models.Text, TranslatedText *models.Text, ClientID uint) (*models.Translation, error) {
	var translation *models.Translation
	err := database.Database.Transaction(func(tx *gorm.DB) error {
		// check source language is supported
		var sourceLanguage *models.Language
		if err := tx.First(&sourceLanguage, SourceText.LanguageID).Error; err != nil {
			return err
		}

		// check translation language is supported
		var translatedLanguage *models.Language
		if err := tx.First(&translatedLanguage, TranslatedText.LanguageID).Error; err != nil {
			return err
		}

		var client *models.Client
		if err := tx.First(&client, ClientID).Error; err != nil {
			return err
		}

		// check whether source text exists
		// create if not
		res := tx.Model(&models.Text{}).
			Where("language_id = ?", SourceText.LanguageID).
			Where("content = ?", SourceText.Content).
			Find(&SourceText)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			if err := tx.Create(&SourceText).Error; err != nil {
				return err
			}
		}

		// check whether translated text exists
		// create if not
		res = tx.Model(&models.Text{}).
			Where("language_id = ?", TranslatedText.LanguageID).
			Where("content = ?", TranslatedText.Content).
			Find(&TranslatedText)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			if err := tx.Create(&TranslatedText).Error; err != nil {
				return err
			}
		}

		// check whether translation exists
		// create if not
		res = tx.Model(&models.Translation{}).
			Where("source_text_id = ?", SourceText.ID).
			Where("target_text_id = ?", TranslatedText.ID).
			Where("client_id = ?", ClientID).
			Find(&translation)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			translation = &models.Translation{
				SourceTextID: SourceText.ID,
				TargetTextID: TranslatedText.ID,
				ClientID: ClientID,
			}
			if err := tx.Create(&translation).Error; err != nil {
				return err
			}
		}
		
		return nil
	})

	if err != nil {
		return nil, err
	}
	
	return translation, nil
}

func Query(SourceLanguageID int64, TargetLanguageID int64, ClientID int64, text string) (*models.Text, error) {
	var translatedText *models.Text
	// search for translations belonging to the client
	res := database.Database.Model(&models.Text{}).
		Joins("left join translations on translations.target_text_id = texts.id").
		Joins("left join texts as source_texts on source_texts.id = translations.source_text_id").
		Where("source_texts.content = ?", text).
		Where("translations.client_id = ?", ClientID).
		First(&translatedText)
	if res.Error == nil {
		return translatedText, nil
	}

	// search for translations from whole pool
	res = database.Database.Model(&models.Text{}).
		Joins("left join translations on translations.target_text_id = texts.id").
		Joins("left join texts as source_texts on source_texts.id = translations.source_text_id").
		Where("source_texts.content = ?", text).
		First(&translatedText)
	if res.Error == nil {
		return translatedText, nil 
	}

	// TODO: fall back on some translation service
	return nil, errors.New("no translation found")
}