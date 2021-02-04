package models

import "github.com/bagus2x/fainmi-be/pkg/entities"

// CreateLinkReq create link request
type CreateLinkReq struct {
	Order   int    `json:"order"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	Display bool   `json:"display"`
}

// CreateLinkRes create link resposne
type CreateLinkRes struct {
	LinkID  int    `json:"linkID"`
	Order   int    `json:"order"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	Display bool   `json:"display"`
}

// GetLinkRes get link response
type GetLinkRes struct {
	LinkID    int    `json:"linkID"`
	ProfileID int    `json:"profileID"`
	Order     int    `json:"order"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	Display   bool   `json:"display"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

// GetLinksRes get links repsonse
type GetLinksRes []*GetLinkRes

// LinkUpdateReq link update request
type LinkUpdateReq struct {
	Order   int    `json:"order"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	Display bool   `json:"display"`
}

// Orders equest struct & response struct
type Orders []*entities.Order
