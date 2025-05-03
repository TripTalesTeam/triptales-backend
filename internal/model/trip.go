package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Trip represents a trip created by a user to a country.
type Trip struct {
	ID      string  `gorm:"primaryKey;type:char(36)" json:"trip_id"`
	UserID      string  `gorm:"type:char(36);not null" json:"user_id"`
	CountryID   string  `gorm:"type:char(36);not null" json:"country_id"`
	Title       string  `gorm:"not null" json:"title"`
	Description string  `gorm:"" json:"description"`
	Latitude    float64 `gorm:"" json:"latitude"`
	Longitude   float64 `gorm:"" json:"longitude"`
	Image       string  `gorm:"" json:"image"`

	User    User    `gorm:"foreignKey:UserID;references:ID" json:"user"`          // Foreign key relation for user
	Country Country `gorm:"foreignKey:CountryID;references:ID" json:"country"` // Foreign key relation for country

	Companions []TripCompanion `gorm:"foreignKey:TripID;constraint:OnDelete:CASCADE;" json:"companions"` // Many-to-many relation with TripCompanion
	Bookmarks  []Bookmark      `gorm:"foreignKey:TripID" json:"bookmarks"`  // One-to-many relation with Bookmark
}

func (t *Trip) BeforeCreate(tx *gorm.DB) error {
	// Generate UUID if not provided
	if t.ID == "" {
		t.ID = uuid.New().String()
	}
	return nil
}
