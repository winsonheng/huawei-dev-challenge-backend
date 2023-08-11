package main

import (
	"backend/database"
	"backend/models"
	"log"

	"github.com/joho/godotenv"
)

func main() {
    loadEnv()
    loadDatabase()
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
