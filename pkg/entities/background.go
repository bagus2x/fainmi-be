package entities

import "database/sql"

// Background struct represent background entity
type Background struct {
	BackgroundID int
	Name         string
	Description  sql.NullString
	CreatedAt    int64
	UpdatedAt    int64
}
