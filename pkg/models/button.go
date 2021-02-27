package models

// CreateButtonRequest -
type CreateButtonRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

// GetButtonResponse -
type GetButtonResponse struct {
	ButtonID    int    `json:"buttonID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

// GetButtonsResponse - -
type GetButtonsResponse []*GetButtonResponse

// UpdateButtonRequest -
type UpdateButtonRequest CreateButtonRequest
