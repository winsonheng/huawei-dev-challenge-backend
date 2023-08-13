package translation

import "backend/models"

type CreateParams struct {
	SourceTextContent     string `json:"text"`
	SourceLanguageID      uint `json:"sourceLanguageID"`
	TranslatedTextContent string `json:"translation"`
	TranslatedLanguageID  uint `json:"targetLanguageID"`
	ClientID              uint   `json:"clientID"`
}

func ToModel(params *CreateParams) (*models.Text, *models.Text, uint) {
	sourceText := models.Text{
		Content: params.SourceTextContent,
		LanguageID: params.SourceLanguageID,
	}

	translatedText := models.Text{
		Content: params.TranslatedTextContent,
		LanguageID: params.TranslatedLanguageID,
	}

	return &sourceText, &translatedText, params.ClientID
}