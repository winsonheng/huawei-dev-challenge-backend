package translation

import (
	"backend/dataaccess/translations"
	"backend/params/translation"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleCreate(context *gin.Context) {
	var createParams translation.CreateParams

	if err := context.ShouldBindJSON(&createParams); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sourceText, translatedText, clientID := translation.ToModel(&createParams)

	translation, err := translations.Create(sourceText, translatedText, clientID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"translation": translation})
}