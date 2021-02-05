package style

import (
	"database/sql"

	"github.com/bagus2x/fainmi-be/pkg/entities"
)

// Repository style
type Repository interface {
	Create(style *entities.Style) error
	Read(profileID int) (*entities.Style, error)
	ReadStyleDetail(profileID string) (*entities.StyleDetail, error)
	Update(profileID int, style *entities.Style) (bool, error)
	Delete(profileID int) (bool, error)
}

type repository struct {
	db *sql.DB
}

// NewRepo -
func NewRepo(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r repository) Create(style *entities.Style) error {
	_, err := r.db.Exec(`INSERT INTO style VALUES($1, $2, $3, $4, $5, $6)`,
		style.ProfileID,
		style.BackgroundID,
		style.ButtonID,
		style.FontID,
		style.CreatedAt,
		style.UpdatedAt,
	)

	return err
}

func (r repository) Read(profileID int) (*entities.Style, error) {
	var style entities.Style
	err := r.db.QueryRow(`SELECT * FROM style WHERE profile_id=$1`, profileID).Scan(
		&style.ProfileID,
		&style.BackgroundID,
		&style.ButtonID,
		&style.FontID,
		&style.CreatedAt,
		&style.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &style, err
}

// ReadStyleDetail -
func (r repository) ReadStyleDetail(username string) (*entities.StyleDetail, error) {
	q := `SELECT profile_id, bgn.name as background, btn.name as button, fnt.name as font FROM style stl
		  LEFT JOIN background bgn USING(background_id)
		  LEFT JOIN button btn USING(button_id)
		  LEFT JOIN font fnt USING(font_id)
		  WHERE profile_id=(SELECT profile_id FROM profile WHERE username=$1)`
	var styleDetail entities.StyleDetail
	err := r.db.QueryRow(q, username).Scan(
		&styleDetail.ProfileID,
		&styleDetail.Background,
		&styleDetail.Button,
		&styleDetail.Font,
	)
	if err != nil {
		return nil, err
	}

	return &styleDetail, nil
}

func (r repository) Update(profileID int, style *entities.Style) (bool, error) {
	q := `UPDATE style SET background_id=$1, button_id=$2, font_id=$3, updated_at=$4 WHERE profile_id=$5`
	res, err := r.db.Exec(
		q,
		style.BackgroundID,
		style.ButtonID,
		style.FontID,
		style.UpdatedAt,
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
	res, err := r.db.Exec(`DELETE FROM style WHERE profile_id=$1`, profileID)
	if err != nil {
		return false, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
