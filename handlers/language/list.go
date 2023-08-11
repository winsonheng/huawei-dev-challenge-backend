package language

import (
	"backend/dataaccess/languages"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleList(context *gin.Context) {
	languages, err := languages.List()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"language": languages})
}