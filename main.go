package main

import (
	"backend/database"
	"backend/models"
	"backend/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    loadEnv()
    loadDatabase()
	serveApplication()
}

func loadDatabase() {
    database.Connect()
    database.Database.AutoMigrate(&models.Client{})
	database.Database.AutoMigrate(&models.Language{})
	database.Database.AutoMigrate(&models.Text{})
	database.Database.AutoMigrate(&models.Translation{})
	database.Database.AutoMigrate(&models.User{})
}

func loadEnv() {
    err := godotenv.Load(".env.local")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func serveApplication() {
    router := gin.Default()

    routes.GetRoutes(router)

    router.Run(":8000")
    fmt.Println("Server running on port 8000")
}