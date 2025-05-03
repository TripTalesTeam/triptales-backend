// Package repository contains repository implementations for data access
package repository

import (
	"github.com/breezjirasak/triptales/internal/model"
	"gorm.io/gorm"
)

// FriendRepository handles DB operations for friendships
type FriendRepository struct {
	DB *gorm.DB
}

// NewFriendRepository creates a new instance of FriendRepository
func NewFriendRepository(db *gorm.DB) *FriendRepository {
	return &FriendRepository{DB: db}
}

// Create adds a new friend relation
func (r *FriendRepository) Create(friend *model.Friend) error {
	return r.DB.Create(friend).Error
}

// FindAllByUserID fetches all friends of a user
func (r *FriendRepository) FindAllByUserID(userID string) ([]model.Friend, error) {
	var friends []model.Friend
	err := r.DB.Preload("Friend").Where("user_id = ?", userID).Find(&friends).Error
	return friends, err
}

// Delete removes a friend relationship
func (r *FriendRepository) Delete(userID, friendID string) error {
	return r.DB.Where("user_id = ? AND friend_id = ?", userID, friendID).Delete(&model.Friend{}).Error
}

// Exists checks if a friend relationship exists
func (r *FriendRepository) Exists(userID, friendID string) (bool, error) {
	var count int64
	err := r.DB.Model(&model.Friend{}).
		Where("user_id = ? AND friend_id = ?", userID, friendID).
		Count(&count).Error
	return count > 0, err
}
