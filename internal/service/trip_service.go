package service

import (
	"errors"
	"github.com/breezjirasak/triptales/internal/model"
	"github.com/breezjirasak/triptales/internal/repository"
)

type TripService struct {
	TripRepo *repository.TripRepository
}

func NewTripService(repo *repository.TripRepository) *TripService {
	return &TripService{TripRepo: repo}
}

func (s *TripService) CreateTrip(trip *model.Trip) error {
	if trip.Title == "" || trip.UserID == "" || trip.CountryID == "" {
		return errors.New("missing required fields")
	}
	return s.TripRepo.Create(trip)
}

func (s *TripService) GetTripByID(id string) (*model.Trip, error) {
	return s.TripRepo.FindByID(id)
}

func (s *TripService) GetAllTrips() ([]model.Trip, error) {
	return s.TripRepo.FindAll()
}

func (s *TripService) UpdateTrip(trip *model.Trip) error {
	return s.TripRepo.Update(trip)
}

func (s *TripService) DeleteTrip(id string) error {
	return s.TripRepo.Delete(id)
}