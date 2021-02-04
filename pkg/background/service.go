package background

import (
	"database/sql"
	"time"

	"github.com/bagus2x/fainmi-be/pkg/entities"
	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/bagus2x/fainmi-be/pkg/models/errors"
)

// Service -
type Service interface {
	AddBackground(req *models.CreateBackgroundReq) error
	GetBackground(BackgroundID int) (*models.GetBackgroundRes, error)
	UpdateBackground(BackgroundID int, req *models.UpdateBackgroundReq) error
	DeleteBackground(BackgroundID int) error
}

type service struct {
	repo Repository
}

// NewService -
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) AddBackground(req *models.CreateBackgroundReq) error {
	background := &entities.Background{
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
		Name:        req.Name,
		Description: sql.NullString{Valid: req.Description != "", String: req.Description},
	}
	err := s.repo.Create(background)
	if err != nil {
		return errors.ErrDatabase(err)
	}

	return nil
}

func (s service) GetBackground(backgroundID int) (*models.GetBackgroundRes, error) {
	background, err := s.repo.Read(backgroundID)
	if err != nil {
		return nil, errors.ErrDatabase(err)
	}

	res := &models.GetBackgroundRes{
		BackgroundID: background.BackgroundID,
		Name:         background.Name,
		Description:  background.Description.String,
		CreatedAt:    background.CreatedAt,
		UpdatedAt:    background.UpdatedAt,
	}

	return res, nil
}

func (s service) UpdateBackground(backgroundID int, req *models.UpdateBackgroundReq) error {
	background := &entities.Background{
		Name:        req.Name,
		Description: sql.NullString{Valid: req.Description != "", String: req.Description},
		UpdatedAt:   time.Now().Unix(),
	}
	isUpdated, err := s.repo.Update(backgroundID, background)
	if err != nil {
		return errors.ErrDatabase(err)
	}
	if !isUpdated {
		return errors.ErrBackgroundNotFound
	}

	return nil
}

func (s service) DeleteBackground(BackgroundID int) error {
	isDeleted, err := s.repo.Delete(BackgroundID)
	if err != nil {
		return errors.ErrDatabase(err)
	}
	if !isDeleted {
		return errors.ErrBackgroundNotFound
	}

	return nil
}
