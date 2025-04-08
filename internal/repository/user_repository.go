package repository

import (
	"github.com/breezjirasak/triptales/config"
	"github.com/breezjirasak/triptales/internal/model"
)

func GetAllUsers() []model.User {
	var users []model.User
	config.DB.Find(&users)
	return users
}

func CreateUser(user model.User) model.User {
	config.DB.Create(&user)
	return user
}