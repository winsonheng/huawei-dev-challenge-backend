package language

import (
	"backend/dataaccess/languages"
	"backend/params/language"
	"backend/views/languageViews"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleCreate(context *gin.Context) {
	var createParams language.CreateParams

	if err := context.ShouldBindJSON(&createParams); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model := language.ToModel(&createParams)

	language, err := languages.Create(model)
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	view := languageViews.ViewFrom(language)

	context.JSON(http.StatusOK, gin.H{"language": view})
}