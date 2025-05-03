package repository

import (
	"errors"
	"github.com/breezjirasak/triptales/internal/model"
	"gorm.io/gorm"
)

// UserRepository handles data access operations for users
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// FindByUsername retrieves a user by username
func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // User not found, but not an error
		}
		return nil, err
	}
	return &user, nil
}

// FindByEmail retrieves a user by email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // User not found, but not an error
		}
		return nil, err
	}
	return &user, nil
}

// FindByID retrieves a user by ID
func (r *UserRepository) FindByID(id string) (*model.User, error) {
	var user model.User
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Create saves a new user to the database
func (r *UserRepository) Create(user *model.User) error {
	return r.DB.Create(user).Error
}

// Update updates an existing user in the database
func (r *UserRepository) Update(user *model.User) error {
	return r.DB.Save(user).Error
}

// UpdateProfileImage updates just the profile image field for a user
func (r *UserRepository) UpdateProfileImage(userID string, imagePath string) error {
	return r.DB.Model(&model.User{}).Where("id = ?", userID).Update("profile_image", imagePath).Error
}