package client

import (
	"backend/dataaccess/clients"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleList(context *gin.Context) {
	clients, err := clients.List()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"clients": clients})
}