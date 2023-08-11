package users

import (
	"backend/database"
	"backend/models"
)

func Create(user *models.User) (*models.User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}