package service

import (
	"errors"

	"github.com/breezjirasak/triptales/internal/model"
	"github.com/breezjirasak/triptales/internal/repository"
)

type FriendService struct {
	Repo *repository.FriendRepository
}

func NewFriendService(repo *repository.FriendRepository) *FriendService {
	return &FriendService{Repo: repo}
}

func (s *FriendService) AddFriend(userID, friendID string) error {
	if userID == friendID {
		return errors.New("cannot add yourself as a friend")
	}
	exists, err := s.Repo.Exists(userID, friendID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("already friends")
	}
	friend := &model.Friend{
		UserID:   userID,
		FriendID: friendID,
	}
	return s.Repo.Create(friend)
}

func (s *FriendService) GetFriends(userID string) ([]model.Friend, error) {
	return s.Repo.FindAllByUserID(userID)
}

func (s *FriendService) RemoveFriend(userID, friendID string) error {
	return s.Repo.Delete(userID, friendID)
}
