// Package repository contains repository implementations for data access
package repository

import (
	"errors"
	"github.com/breezjirasak/triptales/internal/model"
	"gorm.io/gorm"
)

// CountryRepository handles database operations for countries
type CountryRepository struct {
	DB *gorm.DB
}

// NewCountryRepository creates a new instance of CountryRepository
func NewCountryRepository(db *gorm.DB) *CountryRepository {
	return &CountryRepository{DB: db}
}

// Create adds a new country to the database
func (c *CountryRepository) Create(country *model.Country) error {
	return c.DB.Create(country).Error
}

// FindByCountryName retrieves a country by its name
func (c *CountryRepository) FindByCountryName(countryName string) (*model.Country, error) {
	var country model.Country
	if err := c.DB.Where("name = ?", countryName).First(&country).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Country not found, but not an error
		}
		return nil, err
	}
	return &country, nil
}

// FindAll retrieves all countries from the database
func (c *CountryRepository) FindAll() ([]model.Country, error) {
	var countries []model.Country
	if err := c.DB.Find(&countries).Error; err != nil {
		return nil, err
	}
	return countries, nil
}

// Update updates an existing country in the database
func (c *CountryRepository) Update(country *model.Country) error {
	return c.DB.Save(country).Error
}

// Delete removes a country from the database by ID
func (c *CountryRepository) Delete(id string) error {
	return c.DB.Where("id = ?", id).Delete(&model.Country{}).Error
}

// FindByID retrieves a country by its ID
func (c *CountryRepository) FindByID(id string) (*model.Country, error) {
	var country model.Country
	if err := c.DB.Where("id = ?", id).First(&country).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &country, nil
}

