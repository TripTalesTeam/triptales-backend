package handler

import (
	"net/http"
	"github.com/breezjirasak/triptales/internal/model"
	"github.com/breezjirasak/triptales/internal/service"
	"github.com/gin-gonic/gin"
)

type TripHandler struct {
	Service *service.TripService
}

func NewTripHandler(s *service.TripService) *TripHandler {
	return &TripHandler{Service: s}
}

func (h *TripHandler) CreateTrip(c *gin.Context) {
	var trip model.Trip
	if err := c.ShouldBindJSON(&trip); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Auto-inject userID from context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}
	trip.UserID = userID.(string)

	if err := h.Service.CreateTrip(&trip); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, trip)
}


func (h *TripHandler) GetTripByID(c *gin.Context) {
	id := c.Param("id")
	trip, err := h.Service.GetTripByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "trip not found"})
		return
	}
	c.JSON(http.StatusOK, trip)
}

func (h *TripHandler) GetAllTrips(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	userIDStr, ok := userID.(string)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	trips, err := h.Service.GetAllTrips(userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trips)
}

func (h *TripHandler) GetFriendTrip(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Get the country name from query parameters
	countryName := c.Query("country")

	trips, err := h.Service.GetAllFriendTrips(userIDStr, countryName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "trip not found"})
		return
	}

	for i, j := 0, len(trips)-1; i < j; i, j = i+1, j-1 {
		trips[i], trips[j] = trips[j], trips[i]
	}

	c.JSON(http.StatusOK, trips)
}

func (h *TripHandler) GetBookmarkTrip(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Get the country name from query parameters
	countryName := c.Query("country")

	trips, err := h.Service.GetAllBookmarkTrips(userIDStr, countryName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "trip not found"})
		return
	}

	for i, j := 0, len(trips)-1; i < j; i, j = i+1, j-1 {
		trips[i], trips[j] = trips[j], trips[i]
	}

	c.JSON(http.StatusOK, trips)
}

func (h *TripHandler) GetCompanionTrip(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	trips, err := h.Service.GetAllCompanionTrips(userIDStr)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "trip not found"})
		return
	}

	c.JSON(http.StatusOK, trips)
}


func (h *TripHandler) UpdateTrip(c *gin.Context) {
	var trip model.Trip
	if err := c.ShouldBindJSON(&trip); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}
	trip.UserID = userID.(string)
	if err := h.Service.UpdateTrip(&trip); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trip)
}

func (h *TripHandler) DeleteTrip(c *gin.Context) {
	id := c.Param("id")

	// Check if the trip exists first
	trip, err := h.Service.GetTripByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check trip existence"})
		return
	}
	if trip == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "trip not found"})
		return
	}

	// Proceed with deletion
	if err := h.Service.DeleteTrip(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "trip deleted successfully"})
}

