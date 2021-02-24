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
	CreateLink(profileID int, req *models.CreateLinkRequest) (*models.CreateLinkResponse, error)
	GetLink(linkID, profileID int) (*models.GetLinkResponse, error)
	GetLinks(profileID int) (models.GetLinksResponse, error)
	GetPublicLinks(username string) (models.GetLinksResponse, error)
	UpdateLink(linkID, profileID int, req *models.UpdateLinkRequest) (*models.UpdateLinkResponse, error)
	UpdateDisplay(linkID, profileID int, req *models.UpdateDisplayRequest) error
	UpdateLinksOrder(profileID int, order models.LinksOrder) error
	DeleteLink(linkID, profileID int) error
}

type service struct {
	repo Repository
}

// NewService -
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) CreateLink(profileID int, req *models.CreateLinkRequest) (*models.CreateLinkResponse, error) {
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
	res := &models.CreateLinkResponse{
		LinkID:  link.LinkID,
		Order:   link.Order,
		Title:   link.Title.String,
		URL:     link.URL.String,
		Display: link.Display,
	}
	return res, nil
}

func (s service) GetLink(linkID, profileID int) (*models.GetLinkResponse, error) {
	link, err := s.repo.Read(linkID, profileID)
	if err != nil {
		return nil, errors.ErrDatabase(err)
	}

	res := &models.GetLinkResponse{
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

func (s service) GetLinks(profileID int) (models.GetLinksResponse, error) {
	links, err := s.repo.ReadByProfileID(profileID)
	if err != nil {
		return nil, errors.ErrDatabase(err)
	}
	res := make(models.GetLinksResponse, 0)
	for _, link := range links {
		res = append(res, &models.GetLinkResponse{
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

func (s service) GetPublicLinks(username string) (models.GetLinksResponse, error) {
	links, err := s.repo.ReadByUsername(username)
	if err != nil {
		return nil, errors.ErrDatabase(err)
	}
	res := make(models.GetLinksResponse, 0)
	for _, link := range links {
		res = append(res, &models.GetLinkResponse{
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

func (s service) UpdateLink(linkID, profileID int, req *models.UpdateLinkRequest) (*models.UpdateLinkResponse, error) {
	link := &entities.Link{
		Order:     req.Order,
		Title:     sql.NullString{Valid: req.Title != "", String: req.Title},
		URL:       sql.NullString{Valid: req.URL != "", String: req.URL},
		Display:   req.Display,
		UpdatedAt: time.Now().Unix(),
	}
	updated, err := s.repo.Update(linkID, profileID, link)
	if err != nil {
		return nil, errors.ErrDatabase(err)
	}
	if !updated {
		return nil, errors.ErrLinkNotFound
	}

	res := &models.UpdateLinkResponse{
		LinkID:  linkID,
		Display: req.Display,
		Order:   req.Order,
		Title:   req.Title,
		URL:     req.URL,
	}

	return res, nil
}

func (s service) UpdateDisplay(linkID, profileID int, req *models.UpdateDisplayRequest) error {
	updated, err := s.repo.UpdateDisplay(linkID, profileID, req.Display)
	if err != nil {
		return errors.ErrDatabase(err)
	}
	if !updated {
		return errors.ErrLinkNotFound
	}

	return nil
}

func (s service) DeleteLink(linkID, profileID int) error {
	deleted, err := s.repo.Delete(linkID, profileID)
	if err != nil {
		return errors.ErrDatabase(err)
	}
	if !deleted {
		return errors.ErrLinkNotFound
	}

	return nil
}

func (s service) UpdateLinksOrder(profileID int, req models.LinksOrder) error {
	err := s.repo.UpdateOrder(profileID, req)
	if err != nil {
		return errors.ErrDatabase(err)
	}

	return nil
}
