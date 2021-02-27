package models

// CreateFontRequest -
type CreateFontRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	FontFamily  string `json:"fontFamily"`
	Href        string `json:"href"`
}

// GetFontResponse -
type GetFontResponse struct {
	FontID      int    `json:"fontID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	FontFamily  string `json:"fontFamily"`
	Href        string `json:"href"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

// GetFontsResponse -
type GetFontsResponse []*GetFontResponse

// UpdateFontRequest -
type UpdateFontRequest CreateFontRequest
