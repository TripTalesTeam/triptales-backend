package repository

import (
	"github.com/breezjirasak/triptales/internal/model"
	"gorm.io/gorm"
)

type TripCompanionRepository struct {
	DB *gorm.DB
}

func NewTripCompanionRepository(db *gorm.DB) *TripCompanionRepository {
	return &TripCompanionRepository{DB: db}
}

func (r *TripCompanionRepository) Add(companion *model.TripCompanion) error {
	return r.DB.Create(companion).Error
}

func (r *TripCompanionRepository) Remove(tripID, userID string) error {
	return r.DB.Where("trip_id = ? AND user_id = ?", tripID, userID).Delete(&model.TripCompanion{}).Error
}

func (r *TripCompanionRepository) GetCompanionsByTripID(tripID string) ([]model.TripCompanion, error) {
	var companions []model.TripCompanion
	err := r.DB.Preload("User").Preload("Trip").Where("trip_id = ?", tripID).Find(&companions).Error
	return companions, err
}
