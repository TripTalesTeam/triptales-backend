package model

// Friend represents the friendship relationship between users (many-to-many).
type Friend struct {
	FriendID string `gorm:"primaryKey;type:char(36)" json:"friend_id"`
	UserID   string `gorm:"primaryKey;type:char(36)" json:"user_id"`

	User   User `gorm:"foreignKey:UserID;references:ID" json:"user"`     // Foreign key relation for user
	Friend User `gorm:"foreignKey:FriendID;references:ID" json:"friend"` // Foreign key relation for friend
}
