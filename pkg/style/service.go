package style

import (
	"database/sql"
	"time"

	"github.com/bagus2x/fainmi-be/pkg/entities"
	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/bagus2x/fainmi-be/pkg/models/errors"
)

// Service -
type Service interface {
	CreateStyle(profileID int, req *models.StyleRequest) error
	GetStyle(profileID int) (*models.StyleResponse, error)
	GetStyleDetail(username string) (*models.StyleDetail, error)
	UpdateStyle(profileID int, req *models.StyleRequest) error
	DeleteStyle(profileID int) error
}

type service struct {
	repo Repository
}

// NewService -
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) CreateStyle(profileID int, req *models.StyleRequest) error {
	style := &entities.Style{
		ProfileID:    profileID,
		BackgroundID: sql.NullInt32{Valid: req.BackgroundID != 0, Int32: int32(req.BackgroundID)},
		ButtonID:     sql.NullInt32{Valid: req.ButtonID != 0, Int32: int32(req.ButtonID)},
		FontID:       sql.NullInt32{Valid: req.FontID != 0, Int32: int32(req.FontID)},
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}

	err := s.repo.Create(style)
	if err != nil {
		return errors.ErrDatabase(err)
	}

	return nil
}

func (s service) GetStyle(profileID int) (*models.StyleResponse, error) {
	style, err := s.repo.Read(profileID)
	if err != nil {
		return nil, errors.ErrDatabase(err)
	}

	res := &models.StyleResponse{
		ProfileID:    style.ProfileID,
		BackgroundID: int(style.BackgroundID.Int32),
		ButtonID:     int(style.BackgroundID.Int32),
		FontID:       int(style.FontID.Int32),
	}

	return res, nil
}

func (s service) GetStyleDetail(username string) (*models.StyleDetail, error) {
	stl, err := s.repo.ReadStyleDetail(username)
	if err != nil {
		return nil, errors.ErrDatabase(err)
	}

	res := models.StyleDetail{
		ProfileID: stl.ProfileID,
		Background: models.BackgroundStyle{
			Name:     stl.BackgroundName,
			Image:    stl.BackgroundImage,
			SubImage: stl.BackgroundSubImage,
		},
		Button: models.ButtonStyle{
			Name:  stl.ButtonName,
			Image: stl.ButtonImage,
		},
		Font: models.FontStyle{
			Name:       stl.FontName,
			FontFamily: stl.FontFamily,
			Href:       stl.FontHref,
		},
	}

	return &res, nil
}

func (s service) UpdateStyle(profileID int, req *models.StyleRequest) error {
	style := &entities.Style{
		BackgroundID: sql.NullInt32{Valid: req.BackgroundID != 0, Int32: int32(req.BackgroundID)},
		ButtonID:     sql.NullInt32{Valid: req.ButtonID != 0, Int32: int32(req.ButtonID)},
		FontID:       sql.NullInt32{Valid: req.FontID != 0, Int32: int32(req.FontID)},
		UpdatedAt:    time.Now().Unix(),
	}

	updated, err := s.repo.Update(profileID, style)
	if err != nil {
		return err
	}
	if !updated {
		return errors.ErrStyleNotFound
	}

	return nil
}

func (s service) DeleteStyle(profileID int) error {
	deleted, err := s.repo.Delete(profileID)
	if err != nil {
		return err
	}
	if !deleted {
		return errors.ErrStyleNotFound
	}

	return nil
}
