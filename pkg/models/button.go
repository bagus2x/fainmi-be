package models

// CreateButtonReq -
type CreateButtonReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetButtonRes -
type GetButtonRes struct {
	ButtonID    int    `json:"buttonID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

// UpdateButtonReq -
type UpdateButtonReq CreateButtonReq
