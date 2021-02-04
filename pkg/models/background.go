package models

// CreateBackgroundReq -
type CreateBackgroundReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetBackgroundRes -
type GetBackgroundRes struct {
	BackgroundID int    `json:"backgroundID"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
}

// UpdateBackgroundReq -
type UpdateBackgroundReq CreateBackgroundReq
