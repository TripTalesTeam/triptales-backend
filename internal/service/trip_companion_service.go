package service

import (
	"github.com/breezjirasak/triptales/internal/model"
	"github.com/breezjirasak/triptales/internal/repository"
)

type TripCompanionService struct {
	Repo *repository.TripCompanionRepository
}

func NewTripCompanionService(repo *repository.TripCompanionRepository) *TripCompanionService {
	return &TripCompanionService{Repo: repo}
}

func (s *TripCompanionService) AddCompanion(companion *model.TripCompanion) error {
	return s.Repo.Add(companion)
}

func (s *TripCompanionService) RemoveCompanion(tripID, userID string) error {
	return s.Repo.Remove(tripID, userID)
}

func (s *TripCompanionService) GetCompanions(tripID string) ([]model.TripCompanion, error) {
	return s.Repo.GetCompanionsByTripID(tripID)
}
