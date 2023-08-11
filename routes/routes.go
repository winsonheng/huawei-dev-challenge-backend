package routes

import (
	"backend/handlers/language"
	"backend/handlers/user"

	"github.com/gin-gonic/gin"
)

func GetRoutes(r *gin.Engine) {
	getAuthRoutes(r)
	getLanguageRoutes(r)
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