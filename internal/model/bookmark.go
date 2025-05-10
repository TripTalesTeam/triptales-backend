package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Bookmark represents a trip bookmarked by a user (one-to-many).
type Bookmark struct {
	ID string `gorm:"primaryKey;type:char(36)" json:"bookmark_id"`
	UserID     string `gorm:"type:char(36);not null" json:"user_id"`
	TripID     string `gorm:"type:char(36);not null" json:"trip_id"`

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;" json:"user"`
	Trip Trip `gorm:"foreignKey:TripID;references:ID;constraint:OnDelete:CASCADE;" json:"trip"`
}


func (b *Bookmark) BeforeCreate(tx *gorm.DB) error {
	// Generate UUID if not provided
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return nil
}
