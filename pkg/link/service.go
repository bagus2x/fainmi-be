package link

import (
	"database/sql"
	"time"

	"github.com/bagus2x/fainmi-be/pkg/entities"
	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/bagus2x/fainmi-be/pkg/models/errors"
)

// Service -
type Service interface {
	CreateLink(profileID int, req *models.CreateLinkReq) (*models.CreateLinkRes, error)
	GetLink(linkID int) (*models.GetLinkRes, error)
	GetLinks(profileID int, displayAll bool) (models.GetLinksRes, error)
	UpdateLink(profileID int, req *models.LinkUpdateReq) error
	DeleteLink(profileID int) error
	UpdateLinkOrder(orders models.Orders) error
}

type service struct {
	repo Repository
}

// NewService -
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) CreateLink(profileID int, req *models.CreateLinkReq) (*models.CreateLinkRes, error) {
	link := &entities.Link{
		ProfileID: profileID,
		Order:     req.Order,
		Title:     sql.NullString{Valid: req.Title != "", String: req.Title},
		URL:       sql.NullString{Valid: req.URL != "", String: req.URL},
		Display:   req.Display,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	err := s.repo.Create(link)
	if err != nil {
		return nil, errors.ErrDatabase(err)
	}
	res := &models.CreateLinkRes{
		LinkID:  link.LinkID,
		Order:   link.Order,
		Title:   link.Title.String,
		URL:     link.URL.String,
		Display: link.Display,
	}
	return res, nil
}

func (s service) GetLink(linkID int) (*models.GetLinkRes, error) {
	link, err := s.repo.Read(linkID)
	if err != nil {
		return nil, errors.ErrDatabase(err)
	}

	res := &models.GetLinkRes{
		LinkID:    link.LinkID,
		ProfileID: link.ProfileID,
		Order:     link.Order,
		Title:     link.Title.String,
		URL:       link.URL.String,
		Display:   link.Display,
		CreatedAt: link.CreatedAt,
		UpdatedAt: link.UpdatedAt,
	}

	return res, nil
}

func (s service) GetLinks(profileID int, displayAll bool) (models.GetLinksRes, error) {
	links, err := s.repo.ReadByProfileID(profileID, displayAll)
	if err != nil {
		return nil, errors.ErrDatabase(err)
	}
	res := make(models.GetLinksRes, 0)
	for _, link := range links {
		res = append(res, &models.GetLinkRes{
			LinkID:    link.LinkID,
			ProfileID: link.ProfileID,
			Order:     link.Order,
			Title:     link.Title.String,
			URL:       link.URL.String,
			Display:   link.Display,
			CreatedAt: link.CreatedAt,
			UpdatedAt: link.UpdatedAt,
		})
	}

	return res, nil
}

func (s service) UpdateLink(profileID int, req *models.LinkUpdateReq) error {
	link := &entities.Link{
		Order:     req.Order,
		Title:     sql.NullString{Valid: req.Title != "", String: req.Title},
		URL:       sql.NullString{Valid: req.URL != "", String: req.URL},
		Display:   req.Display,
		UpdatedAt: time.Now().Unix(),
	}
	isUpdated, err := s.repo.Update(profileID, link)
	if err != nil {
		return errors.ErrDatabase(err)
	}
	if !isUpdated {
		return errors.ErrLinkNotFound
	}

	return nil
}

func (s service) DeleteLink(linkID int) error {
	isDeleted, err := s.repo.Delete(linkID)
	if err != nil {
		return errors.ErrDatabase(err)
	}
	if !isDeleted {
		return errors.ErrLinkNotFound
	}

	return nil
}

func (s service) UpdateLinkOrder(orders models.Orders) error {
	err := s.repo.UpdateOrder(orders)
	if err != nil {
		return errors.ErrDatabase(err)
	}

	return nil
}
