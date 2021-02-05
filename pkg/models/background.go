package models

// CreateBackgroundRequest -
type CreateBackgroundRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetBackgroundResponse -
type GetBackgroundResponse struct {
	BackgroundID int    `json:"backgroundID"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
}

// GetBackgroundsResponse -
type GetBackgroundsResponse []*GetBackgroundResponse

// UpdateBackgroundRequest -
type UpdateBackgroundRequest CreateBackgroundRequest
