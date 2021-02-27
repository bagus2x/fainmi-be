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
	AddBackground(req *models.CreateBackgroundRequest) error
	GetBackground(BackgroundID int) (*models.GetBackgroundResponse, error)
	GetBackgrounds() (models.GetBackgroundsResponse, error)
	UpdateBackground(BackgroundID int, req *models.UpdateBackgroundRequest) error
	DeleteBackground(BackgroundID int) error
}

type service struct {
	repo Repository
}

// NewService -
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) AddBackground(req *models.CreateBackgroundRequest) error {
	background := &entities.Background{
		Name:        req.Name,
		Description: sql.NullString{Valid: req.Description != "", String: req.Description},
		Image:       req.Image,
		SubImage:    sql.NullString{Valid: req.SubImage != "", String: req.SubImage},
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
	err := s.repo.Create(background)
	if err != nil {
		return errors.ErrDatabase(err)
	}

	return nil
}

func (s service) GetBackground(backgroundID int) (*models.GetBackgroundResponse, error) {
	background, err := s.repo.Read(backgroundID)
	if err != nil {
		return nil, errors.ErrDatabase(err)
	}

	res := &models.GetBackgroundResponse{
		BackgroundID: background.BackgroundID,
		Name:         background.Name,
		Description:  background.Description.String,
		Image:        background.Image,
		SubImage:     background.Description.String,
		CreatedAt:    background.CreatedAt,
		UpdatedAt:    background.UpdatedAt,
	}

	return res, nil
}

func (s service) GetBackgrounds() (models.GetBackgroundsResponse, error) {
	backgrounds, err := s.repo.ReadAll()
	if err != nil {
		return nil, errors.ErrDatabase(err)
	}

	res := make(models.GetBackgroundsResponse, 0)
	for _, background := range backgrounds {
		var bg = models.GetBackgroundResponse{
			BackgroundID: background.BackgroundID,
			Name:         background.Name,
			Description:  background.Description.String,
			Image:        background.Image,
			SubImage:     background.Description.String,
			CreatedAt:    background.CreatedAt,
			UpdatedAt:    background.UpdatedAt,
		}
		res = append(res, &bg)
	}

	return res, nil
}

func (s service) UpdateBackground(backgroundID int, req *models.UpdateBackgroundRequest) error {
	background := &entities.Background{
		Name:        req.Name,
		Description: sql.NullString{Valid: req.Description != "", String: req.Description},
		Image:       req.Image,
		SubImage:    sql.NullString{Valid: req.SubImage != "", String: req.SubImage},
		UpdatedAt:   time.Now().Unix(),
	}

	updated, err := s.repo.Update(backgroundID, background)
	if err != nil {
		return errors.ErrDatabase(err)
	}
	if !updated {
		return errors.ErrBackgroundNotFound
	}

	return nil
}

func (s service) DeleteBackground(BackgroundID int) error {
	deleted, err := s.repo.Delete(BackgroundID)
	if err != nil {
		return errors.ErrDatabase(err)
	}
	if !deleted {
		return errors.ErrBackgroundNotFound
	}

	return nil
}
