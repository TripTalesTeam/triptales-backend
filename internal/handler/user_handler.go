package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/breezjirasak/triptales/internal/model"
	"github.com/breezjirasak/triptales/internal/service"
	"github.com/breezjirasak/triptales/config"
)

func GetUsers(c *gin.Context) {
	users := service.GetUsers()
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // Check for duplicate username or email with a single query
    var existingUser model.User
    result := config.DB.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser)
    
    if result.Error == nil {
        // Found a duplicate, figure out which field matched
        if existingUser.Username == user.Username {
            c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
        } else {
            c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
        }
        return
    }
    
    // If no duplicates, proceed with user creation
    service.AddUser(user)
    c.JSON(http.StatusCreated, gin.H{"success": "Account created"})
}