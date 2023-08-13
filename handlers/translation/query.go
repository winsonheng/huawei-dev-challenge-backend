package translation

import (
	"backend/dataaccess/translations"
	"backend/views/textviews"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleQuery(context *gin.Context) {
	// retrieve query params
	// url should be translations?from={languageID}&to={languageID}&client={clientID}&q={i+want+to+query+something}
	sourceLanguageID, err := strconv.ParseInt(context.Query("from"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	targetLanguageID, err :=  strconv.ParseInt(context.Query("to"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	clientID, err :=  strconv.ParseInt(context.Query("client"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query := context.Query("q")
	text := strings.Join(strings.Split(query, "+"), " ")

	translation, err := translations.Query(sourceLanguageID, targetLanguageID, clientID, text)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	translationView := textviews.ViewFrom(translation)

	context.JSON(http.StatusOK, gin.H{"translation": translationView})
}