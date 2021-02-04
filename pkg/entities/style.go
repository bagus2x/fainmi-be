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
	ProfileID int
	// Background name
	Background sql.NullString
	// Button name
	Button sql.NullString
	// Font name
	Font sql.NullString
}
