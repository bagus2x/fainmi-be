package profile

import (
	"database/sql"

	"github.com/bagus2x/fainmi-be/pkg/entities"
)

// Repository profile
type Repository interface {
	Create(profile *entities.Profile) error
	Read(profileID int) (*entities.Profile, error)
	ReadByUsernameOrEmail(usernameEmail string) (*entities.Profile, error)
	ReadCredentials(username string, email string) (*entities.Credentials, error)
	Update(profileID int, profile *entities.Profile) (bool, error)
	Delete(profileID int) (bool, error)
}

type repository struct {
	db *sql.DB
}

// NewRepo -
func NewRepo(db *sql.DB) Repository {
	return &repository{db: db}
}

// Create method will assign ProfileID
func (r repository) Create(profile *entities.Profile) error {
	q := `INSERT INTO profile VALUES(DEFAULT, $1, $2, $3, $4, $5, $6) RETURNING profile_id`
	err := r.db.QueryRow(
		q,
		profile.Photo,
		profile.Username,
		profile.Email,
		profile.Password,
		profile.CreatedAt,
		profile.UpdatedAt).Scan(&profile.ProfileID)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Read(profileID int) (*entities.Profile, error) {
	var profile entities.Profile
	err := r.db.QueryRow(`SELECT * FROM profile WHERE profile_id=$1`, profileID).Scan(
		&profile.ProfileID,
		&profile.Photo,
		&profile.Username,
		&profile.Email,
		&profile.Password,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &profile, err
}

func (r repository) ReadByUsernameOrEmail(usernameEmail string) (*entities.Profile, error) {
	var profile entities.Profile
	err := r.db.QueryRow(`SELECT * FROM profile WHERE username=$1 OR email=$1`, usernameEmail).Scan(
		&profile.ProfileID,
		&profile.Photo,
		&profile.Username,
		&profile.Email,
		&profile.Password,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &profile, err
}

func (r repository) ReadCredentials(username string, email string) (*entities.Credentials, error) {
	q := `SELECT username, email, password FROM profile WHERE username=$1 OR email=$2`
	var credentials entities.Credentials
	err := r.db.QueryRow(q, username, email).Scan(&credentials.Username, &credentials.Email, &credentials.Password)
	if err != nil {
		return nil, err
	}
	return &credentials, err
}

func (r repository) Update(profileID int, profile *entities.Profile) (bool, error) {
	q := `UPDATE profile SET photo=$1, username=$2, email=$3, password=$4, updated_at=$5 WHERE profile_id=$6`
	res, err := r.db.Exec(
		q,
		profile.Photo,
		profile.Username,
		profile.Email,
		profile.Password,
		profile.UpdatedAt,
		profileID,
	)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (r repository) Delete(profileID int) (bool, error) {
	res, err := r.db.Exec(`DELETE FROM profile WHERE profile_id=$1`, profileID)
	if err != nil {
		return false, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
