package handler

import (
	"net/http"

	"github.com/breezjirasak/triptales/internal/model"
	"github.com/breezjirasak/triptales/internal/service"
	"github.com/gin-gonic/gin"
)

type TripCompanionHandler struct {
	Service *service.TripCompanionService
}

func NewTripCompanionHandler(service *service.TripCompanionService) *TripCompanionHandler {
	return &TripCompanionHandler{Service: service}
}

func (h *TripCompanionHandler) AddCompanion(c *gin.Context) {
	var companion model.TripCompanion
	if err := c.ShouldBindJSON(&companion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.AddCompanion(&companion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "companion added successfully"})
}

func (h *TripCompanionHandler) RemoveCompanion(c *gin.Context) {
	tripID := c.Param("tripId")
	userID := c.Param("userId")

	if err := h.Service.RemoveCompanion(tripID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "companion removed successfully"})
}

func (h *TripCompanionHandler) GetCompanions(c *gin.Context) {
	tripID := c.Param("tripId")

	companions, err := h.Service.GetCompanions(tripID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companions)
}
