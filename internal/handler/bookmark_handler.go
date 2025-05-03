package handler

import (
	"github.com/breezjirasak/triptales/internal/model"
	"github.com/breezjirasak/triptales/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookmarkHandler struct {
	Service *service.BookmarkService
}

func NewBookmarkHandler(service *service.BookmarkService) *BookmarkHandler {
	return &BookmarkHandler{Service: service}
}

func (h *BookmarkHandler) AddBookmark(c *gin.Context) {
	var bookmark model.Bookmark
	if err := c.ShouldBindJSON(&bookmark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}
	bookmark.UserID = userID.(string)

	if err := h.Service.AddBookmark(&bookmark); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Bookmark added successfully"})
}

func (h *BookmarkHandler) RemoveBookmark(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.RemoveBookmark(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bookmark removed successfully"})
}

func (h *BookmarkHandler) GetBookmarks(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	bookmarks, err := h.Service.GetBookmarks(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookmarks)
}
