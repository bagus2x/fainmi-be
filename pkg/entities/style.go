package entities

import "database/sql"

// Style struct represent style entity
type Style struct {
	ProfileID    int
	BackgroundID sql.NullInt32
	ButtonID     sql.NullInt32
	FontID       sql.NullInt32
	CreatedAt    int64
	UpdatedAt    int64
}

// StyleDetail -
type StyleDetail struct {
	ProfileID          int
	BackgroundName     string
	BackgroundImage    string
	BackgroundSubImage string
	ButtonName         string
	ButtonImage        string
	FontName           string
	FontFamily         string
	FontHref           string
}
