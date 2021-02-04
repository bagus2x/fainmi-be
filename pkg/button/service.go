package button

import (
	"database/sql"
	"time"

	"github.com/bagus2x/fainmi-be/pkg/entities"
	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/bagus2x/fainmi-be/pkg/models/errors"
)

// Service -
type Service interface {
	AddButton(req *models.CreateButtonReq) error
	GetButton(buttonID int) (*models.GetButtonRes, error)
	UpdateButton(buttonID int, req *models.UpdateButtonReq) error
	DeleteButton(buttonID int) error
}

type service struct {
	repo Repository
}

// NewService -
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) AddButton(req *models.CreateButtonReq) error {
	button := &entities.Button{
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
		Name:        req.Name,
		Description: sql.NullString{Valid: req.Description != "", String: req.Description},
	}
	err := s.repo.Create(button)
	if err != nil {
		return errors.ErrDatabase(err)
	}

	return nil
}

func (s service) GetButton(buttonID int) (*models.GetButtonRes, error) {
	button, err := s.repo.Read(buttonID)
	if err != nil {
		return nil, errors.ErrDatabase(err)
	}

	res := &models.GetButtonRes{
		ButtonID:    button.ButtonID,
		Name:        button.Name,
		Description: button.Description.String,
		CreatedAt:   button.CreatedAt,
		UpdatedAt:   button.UpdatedAt,
	}

	return res, nil
}

func (s service) UpdateButton(buttonID int, req *models.UpdateButtonReq) error {
	button := &entities.Button{
		Name:        req.Name,
		Description: sql.NullString{Valid: req.Name != "", String: req.Name},
		UpdatedAt:   time.Now().Unix(),
	}
	isUpdated, err := s.repo.Update(buttonID, button)
	if err != nil {
		return errors.ErrDatabase(err)
	}
	if !isUpdated {
		return errors.ErrButtonNotFound
	}

	return nil
}

func (s service) DeleteButton(buttonID int) error {
	isDeleted, err := s.repo.Delete(buttonID)
	if err != nil {
		return errors.ErrDatabase(err)
	}
	if !isDeleted {
		return errors.ErrButtonNotFound
	}

	return nil
}
