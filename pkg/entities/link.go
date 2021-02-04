package entities

import "database/sql"

// Link struct represent link entity
type Link struct {
	LinkID    int
	ProfileID int
	Order     int
	Title     sql.NullString
	URL       sql.NullString
	Display   bool
	CreatedAt int64
	UpdatedAt int64
}

// Order -
type Order struct {
	LinkID int `json:"linkID"`
	Order  int `json:"order"`
}
