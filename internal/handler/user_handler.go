package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/breezjirasak/triptales/internal/service"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
    return &UserHandler{UserService: userService}
}

func (u *UserHandler) GetUsers(c *gin.Context) {
	users, err := u.UserService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve users",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}


func (u *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// Proceed with deletion
	if err := u.UserService.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}


func (u *UserHandler) UpdateUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	var input struct {
		Username     string `json:"username"`
		Email        string `json:"email"`
		ProfileImage string `json:"profile_image"`
	}

	// Bind JSON input to the struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	// Proceed with updating user data
	updatedUser, err := u.UserService.UpdateUser(userID.(string), input.Username, input.Email, input.ProfileImage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update user",
			"details": err.Error(),
		})
		return
	}

	// Return updated user data
	c.JSON(http.StatusOK, updatedUser)
}
