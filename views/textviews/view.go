package textviews

import "backend/models"

type View struct {
	ID         uint   `json:"id"`
	Content    string `json:"content"`
	LangaugeID uint   `json:"languageID"`
}

func ViewFrom(text *models.Text) View {
	return View{
		ID: text.ID,
		Content: text.Content,
		LangaugeID: text.LanguageID,
	}
}