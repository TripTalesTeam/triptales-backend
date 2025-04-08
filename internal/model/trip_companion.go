package model

// TripCompanion represents users who are companions on a trip (many-to-many).
type TripCompanion struct {
	TripID string `gorm:"primaryKey;type:char(36)" json:"trip_id"`
	UserID string `gorm:"primaryKey;type:char(36)" json:"user_id"`	

	Trip Trip `gorm:"foreignKey:TripID;references:ID" json:"trip"` // Foreign key relation for trip
	User User `gorm:"foreignKey:UserID;references:ID" json:"user"` // Foreign key relation for user
}
