// repository/bookmark_repository.go
package repository

import (
	"github.com/breezjirasak/triptales/internal/model"
	"gorm.io/gorm"
)

type BookmarkRepository struct {
	DB *gorm.DB
}

func NewBookmarkRepository(db *gorm.DB) *BookmarkRepository {
	return &BookmarkRepository{DB: db}
}

func (r *BookmarkRepository) Create(bookmark *model.Bookmark) error {
	return r.DB.Create(bookmark).Error
}

func (r *BookmarkRepository) Delete(id string) error {
	return r.DB.Where("trip_id = ?", id).Delete(&model.Bookmark{}).Error
}

func (r *BookmarkRepository) FindByUserID(userID string) ([]model.Bookmark, error) {
	var bookmarks []model.Bookmark
	err := r.DB.Preload("Trip").Where("user_id = ?", userID).Find(&bookmarks).Error
	return bookmarks, err
}