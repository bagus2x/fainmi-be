package button

import (
	"database/sql"

	"github.com/bagus2x/fainmi-be/pkg/entities"
)

// Repository button
type Repository interface {
	Create(button *entities.Button) error
	Read(buttonID int) (*entities.Button, error)
	ReadAll() ([]*entities.Button, error)
	Update(buttonID int, button *entities.Button) (bool, error)
	Delete(buttonID int) (bool, error)
}

type repository struct {
	db *sql.DB
}

// NewRepo -
func NewRepo(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r repository) Create(button *entities.Button) error {
	_, err := r.db.Exec(`INSERT INTO button VALUES(DEFAULT, $1, $2, $3, $4)`,
		button.Name,
		button.Description,
		button.CreatedAt,
		button.UpdatedAt,
	)

	return err
}
func (r repository) Read(buttonID int) (*entities.Button, error) {
	var button entities.Button
	err := r.db.QueryRow(`SELECT * FROM button WHERE button_id=$1`, buttonID).Scan(
		&button.ButtonID,
		&button.Name,
		&button.Description,
		&button.CreatedAt,
		&button.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &button, err
}

func (r repository) ReadAll() ([]*entities.Button, error) {
	rows, err := r.db.Query(`SELECT * FROM button`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	buttons := make([]*entities.Button, 0)

	for rows.Next() {
		var btn entities.Button
		err := rows.Scan(&btn.ButtonID, &btn.Name, &btn.Description, &btn.UpdatedAt, &btn.CreatedAt)
		if err != nil {
			return nil, err
		}
		buttons = append(buttons, &btn)
	}

	return buttons, nil
}

func (r repository) Update(buttonID int, button *entities.Button) (bool, error) {
	res, err := r.db.Exec(`UPDATE button SET name=$1, description=$2, updated_at=$3 WHERE button_id=$4`,
		button.Name,
		button.Description,
		button.UpdatedAt,
		buttonID,
	)
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (r repository) Delete(buttonID int) (bool, error) {
	res, err := r.db.Exec(`DELETE FROM button WHERE button_id=$1`,
		buttonID,
	)
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
