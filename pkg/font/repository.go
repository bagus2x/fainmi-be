package font

import (
	"database/sql"

	"github.com/bagus2x/fainmi-be/pkg/entities"
)

// Repository font
type Repository interface {
	Create(font *entities.Font) error
	Read(fontID int) (*entities.Font, error)
	ReadAll() ([]*entities.Font, error)
	Update(fontID int, font *entities.Font) (bool, error)
	Delete(fontID int) (bool, error)
}

type repository struct {
	db *sql.DB
}

// NewRepo -
func NewRepo(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r repository) Create(font *entities.Font) error {
	_, err := r.db.Exec(`INSERT INTO font VALUES(DEFAULT, $1, $2, $3, $4)`,
		font.Name,
		font.Description,
		font.CreatedAt,
		font.UpdatedAt,
	)

	return err
}
func (r repository) Read(fontID int) (*entities.Font, error) {
	var font entities.Font
	err := r.db.QueryRow(`SELECT * FROM font WHERE font_id=$1`, fontID).Scan(
		&font.FontID,
		&font.Name,
		&font.Description,
		&font.CreatedAt,
		&font.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &font, err
}

func (r repository) ReadAll() ([]*entities.Font, error) {
	rows, err := r.db.Query(`SELECT * FROM font`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fonts := make([]*entities.Font, 0)

	for rows.Next() {
		var font entities.Font
		err := rows.Scan(&font.FontID, &font.Name, &font.Description, &font.UpdatedAt, &font.CreatedAt)
		if err != nil {
			return nil, err
		}
		fonts = append(fonts, &font)
	}

	return fonts, nil
}

func (r repository) Update(fontID int, font *entities.Font) (bool, error) {
	res, err := r.db.Exec(`UPDATE font SET name=$1, description=$2, updated_at=$3 WHERE font_id=$4`,
		font.Name,
		font.Description,
		font.UpdatedAt,
		fontID,
	)

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (r repository) Delete(fontID int) (bool, error) {
	res, err := r.db.Exec(`DELETE FROM font WHERE font_id=$1`,
		fontID,
	)

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
