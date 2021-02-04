package entities

import "database/sql"

// Button struct represent background entity
type Button struct {
	ButtonID    int
	Name        string
	Description sql.NullString
	CreatedAt   int64
	UpdatedAt   int64
}
