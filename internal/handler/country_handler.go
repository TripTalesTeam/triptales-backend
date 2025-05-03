package handler

import (
	"net/http"

	"github.com/breezjirasak/triptales/internal/model"
	"github.com/breezjirasak/triptales/internal/service"
	"github.com/gin-gonic/gin"
)

type CountryHandler struct {
	Service *service.CountryService
}

func NewCountryHandler(service *service.CountryService) *CountryHandler {
	return &CountryHandler{Service: service}
}

func (h *CountryHandler) CreateCountry(c *gin.Context) {
	var country model.Country
	if err := c.ShouldBindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateCountry(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, country)
}

func (h *CountryHandler) GetCountryByID(c *gin.Context) {
	id := c.Param("id")
	country, err := h.Service.GetCountryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, country)
}

func (h *CountryHandler) GetCountryByName(c *gin.Context) {
	name := c.Query("name")
	country, err := h.Service.GetCountryByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, country)
}

func (h *CountryHandler) GetAllCountries(c *gin.Context) {
	countries, err := h.Service.GetAllCountries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, countries)
}

func (h *CountryHandler) UpdateCountry(c *gin.Context) {
	var country model.Country
	if err := c.ShouldBindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.UpdateCountry(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, country)
}

func (h *CountryHandler) DeleteCountry(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.DeleteCountry(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "country deleted successfully"})
}
