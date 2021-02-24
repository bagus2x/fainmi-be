package models

import "github.com/bagus2x/fainmi-be/pkg/entities"

// CreateLinkRequest create link request
type CreateLinkRequest struct {
	Order   int    `json:"order"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	Display bool   `json:"display"`
}

// CreateLinkResponse create link resposne
type CreateLinkResponse struct {
	LinkID  int    `json:"linkID"`
	Order   int    `json:"order"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	Display bool   `json:"display"`
}

// GetLinkResponse get link response
type GetLinkResponse struct {
	LinkID    int    `json:"linkID"`
	ProfileID int    `json:"profileID"`
	Order     int    `json:"order"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	Display   bool   `json:"display"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

// GetLinksResponse -
type GetLinksResponse []*GetLinkResponse

// UpdateLinkRequest -
type UpdateLinkRequest CreateLinkRequest

// UpdateDisplayRequest -
type UpdateDisplayRequest struct {
	Display bool `json:"display"`
}

// UpdateLinkResponse -
type UpdateLinkResponse CreateLinkResponse

// LinksOrder -
type LinksOrder []*entities.Order
