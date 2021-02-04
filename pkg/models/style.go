package models

// StyleReq -
type StyleReq struct {
	BackgroundID int `json:"background_id"`
	ButtonID     int `json:"button_id"`
	FontID       int `json:"font_id"`
}

// StyleRes -
type StyleRes struct {
	ProfileID    int `json:"profile_id"`
	BackgroundID int `json:"background_id"`
	ButtonID     int `json:"button_id"`
	FontID       int `json:"font_id"`
}

// StyleDetail -
type StyleDetail struct {
	ProfileID int
	// Background name
	Background string
	// Button name
	Button string
	// Font name
	Font string
}
