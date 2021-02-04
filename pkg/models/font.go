package models

// CreateFontReq -
type CreateFontReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetFontRes -
type GetFontRes struct {
	FontID      int    `json:"fontID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

// UpdateFontReq -
type UpdateFontReq CreateFontReq
