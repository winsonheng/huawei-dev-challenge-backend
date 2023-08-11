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

func FindUserByUsername(username string) (*models.User, error) {
    var user *models.User
    err := database.Database.Where("username=?", username).Find(&user).Error
    if err != nil {
        return nil, err
    }
    return user, nil
}