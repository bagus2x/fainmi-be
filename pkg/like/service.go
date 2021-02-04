package like

import (
	"time"

	"github.com/bagus2x/fainmi-be/pkg/entities"
	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/bagus2x/fainmi-be/pkg/models/errors"
)

// Service -
type Service interface {
	AddLike(likerID int, like *models.AddLikeReq) error
	GetLikeDetail(linkID int) (*models.LikeDetailRes, error)
	DeleteLike(linkID, likerID int) error
	GetNumberOfLikes(linkID int) (int, error)
}

type service struct {
	repo Repository
}

// NewService -
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// AddLike -
func (s service) AddLike(likerID int, req *models.AddLikeReq) error {
	like := &entities.Like{
		LinkID:    req.LinkID,
		LikerID:   likerID,
		OwnerID:   req.OwnerID,
		CreatedAt: time.Now().Unix(),
	}

	err := s.repo.Create(like)
	if err != nil {
		return errors.ErrDatabase(err)
	}

	return nil
}

func (s service) GetLikeDetail(linkID int) (*models.LikeDetailRes, error) {
	like, err := s.repo.Read(linkID)
	if err != nil {
		return nil, errors.ErrDatabase(err)
	}

	res := &models.LikeDetailRes{
		LikerID:   like.LikerID,
		CreatedAt: like.CreatedAt,
		LinkID:    like.LinkID,
		OwnerID:   like.OwnerID,
	}

	return res, nil
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

func (s service) GetNumberOfLikes(linkID int) (int, error) {
	n, err := s.repo.CountNumberOfLikes(linkID)
	if err != nil {
		return -1, errors.ErrDatabase(err)
	}

	return n, nil
}
