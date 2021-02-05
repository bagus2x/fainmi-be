package background

import (
	"database/sql"

	"github.com/bagus2x/fainmi-be/pkg/entities"
)

// Repository background
type Repository interface {
	Create(background *entities.Background) error
	Read(backgroundID int) (*entities.Background, error)
	Update(backgroundID int, background *entities.Background) (bool, error)
	Delete(backgroundID int) (bool, error)
	ReadAll() ([]*entities.Background, error)
}

type repository struct {
	db *sql.DB
}

// NewRepo -
func NewRepo(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r repository) Create(background *entities.Background) error {
	_, err := r.db.Exec(`INSERT INTO background VALUES(DEFAULT, $1, $2, $3, $4)`,
		background.Name,
		background.Description,
		background.CreatedAt,
		background.UpdatedAt,
	)

	return err
}

func (r repository) Read(backgroundID int) (*entities.Background, error) {
	var background entities.Background
	err := r.db.QueryRow(`SELECT * FROM background WHERE background_id=$1`, backgroundID).Scan(
		&background.BackgroundID,
		&background.Name,
		&background.Description,
		&background.CreatedAt,
		&background.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &background, err
}

func (r repository) ReadAll() ([]*entities.Background, error) {
	rows, err := r.db.Query(`SELECT * FROM background`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	backgrounds := make([]*entities.Background, 0)

	for rows.Next() {
		var bg entities.Background
		err := rows.Scan(&bg.BackgroundID, &bg.Name, &bg.Description, &bg.UpdatedAt, &bg.CreatedAt)
		if err != nil {
			return nil, err
		}
		backgrounds = append(backgrounds, &bg)
	}

	return backgrounds, nil
}

func (r repository) Update(backgroundID int, background *entities.Background) (bool, error) {
	res, err := r.db.Exec(`UPDATE background SET name=$1, description=$2, updated_at=$3 WHERE background_id=$4`,
		background.Name,
		background.Description,
		background.UpdatedAt,
		backgroundID,
	)
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (r repository) Delete(backgroundID int) (bool, error) {
	res, err := r.db.Exec(`DELETE FROM background WHERE background_id=$1`,
		backgroundID,
	)
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
