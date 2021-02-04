package entities

import "database/sql"

// Profile struct represent profile entity
type Profile struct {
	ProfileID int
	Photo     sql.NullString
	Username  string
	Email     string
	Password  string
	CreatedAt int64
	UpdatedAt int64
}

// Credentials -
type Credentials struct {
	Username string
	Email    string
	Password string
}
