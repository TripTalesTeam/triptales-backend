package service

import (
	"github.com/breezjirasak/triptales/internal/model"
	"github.com/breezjirasak/triptales/internal/repository"
)

func GetUsers() []model.User {
	return repository.GetAllUsers()
}

func AddUser(user model.User) model.User {
	return repository.CreateUser(user)
}
