package entities

import "database/sql"

// Font struct represent font entity
type Font struct {
	FontID      int
	Name        string
	Description sql.NullString
	FontFamily  string
	Href        string
	CreatedAt   int64
	UpdatedAt   int64
}
