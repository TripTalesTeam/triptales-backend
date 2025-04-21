package handler

import (
	"net/http"
	"path/filepath"
	"github.com/breezjirasak/triptales/internal/service"
	"github.com/breezjirasak/triptales/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

// NewAuthHandler creates a new instance of AuthHandler
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

// Register handles user registration
func (h *AuthHandler) Register(c *gin.Context) {
	var req service.RegisterRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Check if passwords match
	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "passwords do not match"})
		return
	}
	
	response, err := h.AuthService.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, response)
}

// Login handles user authentication
func (h *AuthHandler) Login(c *gin.Context) {
	var req service.LoginRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	response, err := h.AuthService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, response)
}

// UploadProfileImage handles profile image uploads
func (h *AuthHandler) UploadProfileImage(c *gin.Context) {
	// Extract user ID from JWT claim
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	
	// Get the file from form data
	file, err := c.FormFile("profile_image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no file uploaded"})
		return
	}
	
	// Check file extension
	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "only jpg, jpeg, png, and gif images are allowed"})
		return
	}
	
	// Generate a unique filename
	filename := uuid.New().String() + ext
	
	// Save the file
	dst := filepath.Join("./uploads/profiles", filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}
	
	// Update the user's profile image in the database
	imagePath, err := h.AuthService.UploadProfileImage(userID.(string), filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update profile image"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"profile_image": imagePath})
}

// GetMe returns the current user's information
func (h *AuthHandler) GetMe(c *gin.Context) {
	// Extract user ID from JWT claim
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	
	user, err := h.AuthService.GetUserByID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user information"})
		return
	}
	
	c.JSON(http.StatusOK, user)
}

// RegisterRoutes registers auth routes to the provided router
func (h *AuthHandler) RegisterRoutes(router *gin.Engine) {
	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/register", h.Register)
		authGroup.POST("/login", h.Login)
		
		// Protected routes
		authGroup.Use(middleware.JWTMiddleware())
		{
			authGroup.GET("/me", h.GetMe)
			authGroup.POST("/profile-image", h.UploadProfileImage)
		}
	}
}