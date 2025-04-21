package service

import (
	"errors"
	"path/filepath"
	"time"
	"github.com/breezjirasak/triptales/internal/model" // Replace with your actual model package path
	"github.com/breezjirasak/triptales/internal/auth"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

type RegisterRequest struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	Email           string `json:"email" binding:"required,email"`
	ProfileImage    string `json:"profile_image"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token    string      `json:"token"`
	User     model.User  `json:"user"`
	ExpireAt time.Time   `json:"expire_at"`
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

// Register creates a new user account
func (s *AuthService) Register(req RegisterRequest) (*AuthResponse, error) {
	// Check if username exists
	var existingUser model.User
	if result := s.DB.Where("username = ?", req.Username).First(&existingUser); result.RowsAffected > 0 {
		return nil, errors.New("username already exists")
	}

	// Check if email exists
	if result := s.DB.Where("email = ?", req.Email).First(&existingUser); result.RowsAffected > 0 {
		return nil, errors.New("email already exists")
	}

	// Create new user
	user := model.User{
		ID:           uuid.New().String(),
		Username:     req.Username,
		Password:     req.Password, // Password will be hashed by BeforeCreate hook
		Email:        req.Email,
		ProfileImage: req.ProfileImage,
	}

	if err := s.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, err
	}

	// User object to return (without password)
	user.Password = "" // Clear password for security

	return &AuthResponse{
		Token:    token,
		User:     user,
		ExpireAt: time.Now().Add(24 * time.Hour),
	}, nil
}

// Login authenticates a user and returns a token
func (s *AuthService) Login(req LoginRequest) (*AuthResponse, error) {
	// Find user by username
	var user model.User
	if result := s.DB.Where("username = ?", req.Username).First(&user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid username or password")
		}
		return nil, result.Error
	}

	// Check password
	if !user.CheckPassword(req.Password) {
		return nil, errors.New("invalid username or password")
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, err
	}

	// User object to return (without password)
	user.Password = "" // Clear password for security

	return &AuthResponse{
		Token:    token,
		User:     user,
		ExpireAt: time.Now().Add(24 * time.Hour),
	}, nil
}

// UploadProfileImage handles the profile image upload
func (s *AuthService) UploadProfileImage(userID string, filename string) (string, error) {
	// In a production app, you'd save the file to a storage service and return the URL
	// For simplicity, we'll just assume it's stored locally in an uploads directory
	
	// Generate a unique filename to prevent collisions
	uniqueFilename := uuid.New().String() + filepath.Ext(filename)
	imagePath := "/uploads/profiles/" + uniqueFilename
	
	// Update the user's profile image in the database
	if err := s.DB.Model(&model.User{}).Where("id = ?", userID).Update("profile_image", imagePath).Error; err != nil {
		return "", err
	}
	
	return imagePath, nil
}

// GetUserByID retrieves user details by ID
func (s *AuthService) GetUserByID(userID string) (*model.User, error) {
	var user model.User
	if err := s.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	
	// Remove sensitive information
	user.Password = ""
	
	return &user, nil
}