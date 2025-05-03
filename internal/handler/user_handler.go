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
