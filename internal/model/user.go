package model

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents a user.
type User struct {
	ID           string `gorm:"primaryKey;type:char(36)" json:"user_id"`
	Username     string `gorm:"unique;not null" json:"username"`
	Password     string `gorm:"not null" json:"password"` // Using json:"-" to prevent password from appearing in JSON responses
	Email        string `gorm:"unique;not null" json:"email"`
	ProfileImage string `json:"profile_image"`

	Trips      []Trip          `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"trips"` // One-to-many relation with Trip
	Friends    []Friend        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"friends"`                            // Many-to-many relation with Friend
	Companions []TripCompanion `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"companions"`                         // Many-to-many relation with TripCompanion
	Bookmarks  []Bookmark      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"bookmarks"`                          // One-to-many relation with Bookmark
}

// BeforeCreate will set a UUID rather than numeric ID and hash the password
func (u *User) BeforeCreate(tx *gorm.DB) error {
	// Generate UUID if not provided
	if u.ID == "" {
		u.ID = uuid.New().String()
	}

	// Hash password if not already hashed
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}

	return nil
}

// CheckPassword verifies if the provided password matches the hashed password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
