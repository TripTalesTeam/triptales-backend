package service

import (
	"github.com/breezjirasak/triptales/internal/model"
	"github.com/breezjirasak/triptales/internal/repository"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService{
	return &UserService{UserRepo: userRepo}

}

func (u *UserService) GetUsers() ([]model.User, error) {
	return u.UserRepo.GetAllUsers()
}

func (u *UserService) DeleteUser(id string) error {
	return u.UserRepo.Delete(id)
}
