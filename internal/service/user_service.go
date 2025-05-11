package service

import (
	"errors"

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


func (u *UserService) UpdateUser(id, username, email, profileImage string) (*model.User, error) {
	// Get the user by ID from repository
	user, err := u.UserRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Update the fields
	user.Username = username
	user.Email = email
	user.ProfileImage = profileImage

	// Save the updated user back to the repository
	if err := u.UserRepo.Update(user); err != nil {
		return nil, errors.New("failed to update user")
	}

	return user, nil
}