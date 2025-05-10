package handler

import (
	"net/http"

	"github.com/breezjirasak/triptales/internal/service"
	"github.com/gin-gonic/gin"
)

type FriendHandler struct {
	Service *service.FriendService
}

func NewFriendHandler(service *service.FriendService) *FriendHandler {
	return &FriendHandler{Service: service}
}

// POST /api/friends
func (h *FriendHandler) AddFriend(c *gin.Context) {
	var req struct {
		FriendID string `json:"friend_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userID := c.GetString("userID")
	if err := h.Service.AddFriend(userID, req.FriendID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.AddFriend(req.FriendID, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Friend added"})
}

// GET /api/friends
func (h *FriendHandler) GetFriends(c *gin.Context) {
	userID := c.GetString("userID")
	friends, err := h.Service.GetFriends(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, friends)
}

// DELETE /api/friends/:friend_id
func (h *FriendHandler) RemoveFriend(c *gin.Context) {
	friendID := c.Param("friend_id")
	userID := c.GetString("userID")
	if err := h.Service.RemoveFriend(userID, friendID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Friend removed"})
}
