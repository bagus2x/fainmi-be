package like

import (
	"time"

	"github.com/bagus2x/fainmi-be/pkg/entities"
	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/bagus2x/fainmi-be/pkg/models/errors"
)

// Service -
type Service interface {
	AddLike(linkID, likerID int) error
	GetLikes(linkID int) (models.LikesDetailResponse, error)
	GetNumberOfLikes(linkID int) (int, error)
	DeleteLike(linkID, likerID int) error
}

type service struct {
	repo Repository
}

// NewService -
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// AddLike -
func (s service) AddLike(linkID, likerID int) error {
	like := &entities.Like{
		LinkID:    linkID,
		LikerID:   likerID,
		CreatedAt: time.Now().Unix(),
	}

	err := s.repo.Create(like)
	if err != nil {
		return errors.ErrDatabase(err)
	}

	return nil
}

func (s service) GetLikes(linkID int) (models.LikesDetailResponse, error) {
	res, err := s.repo.Read(linkID)
	if err != nil {
		return nil, errors.ErrDatabase(err)
	}

	return res, nil
}

func (s service) GetNumberOfLikes(linkID int) (int, error) {
	n, err := s.repo.CountNumberOfLikes(linkID)
	if err != nil {
		return -1, errors.ErrDatabase(err)
	}

	return n, nil
}

func (s service) DeleteLike(linkID, likerID int) error {
	isDeleted, err := s.repo.Delete(linkID, likerID)
	if err != nil {
		return errors.ErrDatabase(err)
	}
	if !isDeleted {
		return errors.ErrLikeNotFound
	}

	return nil
}
