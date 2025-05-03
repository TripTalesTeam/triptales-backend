package repository

import (
	"github.com/breezjirasak/triptales/internal/model"
	"gorm.io/gorm"
)

type TripRepository struct {
	DB *gorm.DB
}

func NewTripRepository(db *gorm.DB) *TripRepository {
	return &TripRepository{DB: db}
}

func (r *TripRepository) Create(trip *model.Trip) error {
	return r.DB.Create(trip).Error
}

func (r *TripRepository) FindByID(id string) (*model.Trip, error) {
	var trip model.Trip
	err := r.DB.Preload("User").Preload("Country").Preload("Companions").Preload("Bookmarks").First(&trip, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &trip, nil
}

func (r *TripRepository) FindAll() ([]model.Trip, error) {
	var trips []model.Trip
	err := r.DB.Preload("User").Preload("Country").Preload("Companions").Preload("Bookmarks").Find(&trips).Error
	return trips, err
}

func (r *TripRepository) Update(trip *model.Trip) error {
	return r.DB.Save(trip).Error
}

func (r *TripRepository) Delete(id string) error {
	return r.DB.Delete(&model.Trip{}, "id = ?", id).Error
}