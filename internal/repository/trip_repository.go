package repository

import (
	"github.com/breezjirasak/triptales/internal/model"
	"gorm.io/gorm"
	"strings"
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

func (r *TripRepository) FindByFriendTrips(userId string, countryName string) ([]model.Trip, error) {
	var trips []model.Trip
	db := r.DB.
		Preload("User").
		Preload("Country").
		Preload("Companions").
		Joins("JOIN friends ON trips.user_id = friends.friend_id").
		Where("friends.user_id = ?", userId)

	if countryName != "" {
		db = db.Joins("JOIN countries ON trips.country_id = countries.id").
			Where("LOWER(countries.name) LIKE ?", "%"+strings.ToLower(countryName)+"%")
	}

	err := db.Find(&trips).Error
	return trips, err
}

func (r *TripRepository) FindByBookmarkTrips(userId string, countryName string) ([]model.Trip, error) {
	var trips []model.Trip
	db := r.DB.
		Preload("User").
		Preload("Country").
		Preload("Companions").
		Joins("JOIN bookmarks ON trips.trip_id = bookmarks.trip_id").
		Where("bookmarks.user_id = ?", userId)

	if countryName != "" {
		db = db.Joins("JOIN countries ON trips.country_id = countries.id").
			Where("LOWER(countries.name) LIKE ?", "%"+strings.ToLower(countryName)+"%")
	}

	err := db.Find(&trips).Error
	return trips, err
}
