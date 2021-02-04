package font

import (
	"database/sql"
	"time"

	"github.com/bagus2x/fainmi-be/pkg/entities"
	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/bagus2x/fainmi-be/pkg/models/errors"
)

// Service -
type Service interface {
	AddFont(req *models.CreateFontReq) error
	GetFont(fontID int) (*models.GetFontRes, error)
	UpdateFont(fontID int, req *models.UpdateFontReq) error
	DeleteFont(fontID int) error
}

type service struct {
	repo Repository
}

// NewService -
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) AddFont(req *models.CreateFontReq) error {
	font := &entities.Font{
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
		Name:        req.Name,
		Description: sql.NullString{Valid: req.Description != "", String: req.Description},
	}
	err := s.repo.Create(font)
	if err != nil {
		return errors.ErrDatabase(err)
	}

	return nil
}

func (s service) GetFont(fontID int) (*models.GetFontRes, error) {
	font, err := s.repo.Read(fontID)
	if err != nil {
		return nil, errors.ErrDatabase(err)
	}

	res := &models.GetFontRes{
		FontID:      font.FontID,
		Name:        font.Name,
		Description: font.Description.String,
		CreatedAt:   font.CreatedAt,
		UpdatedAt:   font.UpdatedAt,
	}

	return res, nil
}

func (s service) UpdateFont(fontID int, req *models.UpdateFontReq) error {
	font := &entities.Font{
		Name:        req.Name,
		Description: sql.NullString{Valid: req.Name != "", String: req.Name},
		UpdatedAt:   time.Now().Unix(),
	}
	isUpdated, err := s.repo.Update(fontID, font)
	if err != nil {
		return errors.ErrDatabase(err)
	}
	if !isUpdated {
		return errors.ErrFontNotFound
	}

	return nil
}

func (s service) DeleteFont(fontID int) error {
	isDeleted, err := s.repo.Delete(fontID)
	if err != nil {
		return errors.ErrDatabase(err)
	}
	if !isDeleted {
		return errors.ErrFontNotFound
	}

	return nil
}