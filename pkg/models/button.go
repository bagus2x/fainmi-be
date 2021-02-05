package models

// CreateButtonRequest -
type CreateButtonRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetButtonResponse -
type GetButtonResponse struct {
	ButtonID    int    `json:"buttonID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

// GetButtonsResponse - -
type GetButtonsResponse []*GetButtonResponse

// UpdateButtonRequest -
type UpdateButtonRequest CreateButtonRequest
