package routes

import (
	"backend/handlers/client"
	"backend/handlers/language"
	"backend/handlers/translation"
	"backend/handlers/user"

	"github.com/gin-gonic/gin"
)

func GetRoutes(r *gin.Engine) {
	getAuthRoutes(r)
	getLanguageRoutes(r)
	getClientRoutes(r)
	getTranslationRoutes(r)
}

func getAuthRoutes(r *gin.Engine) {
	authRoutes := r.Group("/auth")
	authRoutes.POST("/register", user.Register)
	authRoutes.POST("/login", user.Login)
}

func getLanguageRoutes(r *gin.Engine) {
	languageRoutes := r.Group("/languages")
	languageRoutes.GET("", language.HandleList)
	languageRoutes.POST("", language.HandleCreate)
}

func getClientRoutes(r *gin.Engine) {
	clientRoutes := r.Group("/clients")
	clientRoutes.GET("", client.HandleList)
	clientRoutes.POST("", client.HandleCreate)
}

func getTranslationRoutes(r *gin.Engine) {
	translationRoutes := r.Group("/translations")
	translationRoutes.POST("", translation.HandleCreate)
}