package service

import (
	"github.com/breezjirasak/triptales/internal/model"
	"github.com/breezjirasak/triptales/internal/repository"
)

type BookmarkService struct {
	Repo *repository.BookmarkRepository
}

func NewBookmarkService(repo *repository.BookmarkRepository) *BookmarkService {
	return &BookmarkService{Repo: repo}
}

func (s *BookmarkService) AddBookmark(bookmark *model.Bookmark) error {
	return s.Repo.Create(bookmark)
}

func (s *BookmarkService) RemoveBookmark(id string) error {
	return s.Repo.Delete(id)
}

func (s *BookmarkService) GetBookmarks(userID string) ([]model.Bookmark, error) {
	return s.Repo.FindByUserID(userID)
}
