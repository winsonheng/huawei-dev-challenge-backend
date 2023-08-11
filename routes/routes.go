package routes

import (
	"backend/handlers/user"

	"github.com/gin-gonic/gin"
)

func GetRoutes(r *gin.Engine) {
	getAuthRoutes(r)
}

func getAuthRoutes(r *gin.Engine) {
	authRoutes := r.Group("/auth")
	authRoutes.POST("/register", user.Register)
	authRoutes.POST("/login", user.Login)
}