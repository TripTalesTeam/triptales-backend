package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Country represents a country.
type Country struct {
	ID    string `gorm:"primaryKey;type:char(36)" json:"country_id"`
	Name         string `gorm:"unique;not null" json:"name"`
	CountryImage string `gorm:"not null" json:"country_image"`

	Trips []Trip `gorm:"foreignKey:CountryID" json:"trips"` // One-to-many relation with Trip
}

func (c *Country) BeforeCreate(tx *gorm.DB) error {
	// Generate UUID if not provided
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}
