package service

import (
	"errors"
	"path/filepath"
	"time"
	"github.com/breezjirasak/triptales/internal/model"
	"github.com/breezjirasak/triptales/internal/auth"
	"github.com/breezjirasak/triptales/internal/repository"
	"github.com/google/uuid"
)

type AuthService struct {
	UserRepo *repository.UserRepository
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
func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

// Register creates a new user account
func (s *AuthService) Register(req RegisterRequest) (*AuthResponse, error) {
	// Check if username exists
	existingUser, err := s.UserRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	// Check if email exists
	existingUser, err = s.UserRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
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

	if err := s.UserRepo.Create(&user); err != nil {
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
	user, err := s.UserRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid username or password")
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
		User:     *user,
		ExpireAt: time.Now().Add(24 * time.Hour),
	}, nil
}

// UploadProfileImage handles the profile image upload
func (s *AuthService) UploadProfileImage(userID string, filename string) (string, error) {
	// Generate a unique filename to prevent collisions
	uniqueFilename := uuid.New().String() + filepath.Ext(filename)
	imagePath := "/uploads/profiles/" + uniqueFilename
	
	// Update the user's profile image in the database
	if err := s.UserRepo.UpdateProfileImage(userID, imagePath); err != nil {
		return "", err
	}
	
	return imagePath, nil
}

// GetUserByID retrieves user details by ID
func (s *AuthService) GetUserByID(userID string) (*model.User, error) {
	user, err := s.UserRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	
	// Remove sensitive information
	user.Password = ""
	
	return user, nil
}