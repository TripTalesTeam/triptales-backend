// Package service contains business logic implementations
package service

import (
	"errors"
	"github.com/breezjirasak/triptales/internal/model"
	"github.com/breezjirasak/triptales/internal/repository"
)

// CountryService handles business logic for country-related operations
type CountryService struct {
	CountryRepo *repository.CountryRepository
}

// NewCountryService creates a new instance of CountryService
func NewCountryService(countryRepo *repository.CountryRepository) *CountryService {
	return &CountryService{CountryRepo: countryRepo}
}

// CreateCountry creates a new country
func (s *CountryService) CreateCountry(country *model.Country) error {
	// Validate input
	if country.Name == "" {
		return errors.New("country name cannot be empty")
	}
	
	// Check if country with the same name already exists
	existing, err := s.CountryRepo.FindByCountryName(country.Name)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("country with this name already exists")
	}
	
	// Create country in database
	return s.CountryRepo.Create(country)
}

// GetCountryByName retrieves a country by its name
func (s *CountryService) GetCountryByName(name string) (*model.Country, error) {
	return s.CountryRepo.FindByCountryName(name)
}

// GetAllCountries retrieves all countries
func (s *CountryService) GetAllCountries() ([]model.Country, error) {
	return s.CountryRepo.FindAll()
}

// UpdateCountry updates an existing country
func (s *CountryService) UpdateCountry(country *model.Country) error {
	// Validate input
	if country.Name == "" {
		return errors.New("country name cannot be empty")
	}
	
	// Check if country exists
	existing, err := s.CountryRepo.FindByID(country.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("country not found")
	}
	
	// Check if new name already exists (only if name is changing)
	if country.Name != existing.Name {
		countryWithSameName, err := s.CountryRepo.FindByCountryName(country.Name)
		if err != nil {
			return err
		}
		
		if countryWithSameName != nil && countryWithSameName.ID != country.ID {
			return errors.New("country with this name already exists")
		}
	}
	
	// Update country in database
	return s.CountryRepo.Update(country)
}

// DeleteCountry removes a country by ID
func (s *CountryService) DeleteCountry(id string) error {
	return s.CountryRepo.Delete(id)
}

// GetCountryByID retrieves a country by its ID
func (s *CountryService) GetCountryByID(id string) (*model.Country, error) {
	return s.CountryRepo.FindByID(id)
}