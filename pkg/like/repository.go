package like

import (
	"database/sql"

	"github.com/bagus2x/fainmi-be/pkg/entities"
)

// Repository like
type Repository interface {
	Create(like *entities.Like) error
	Read(linkID int) ([]*entities.Like, error)
	Delete(linkID, likerID int) (bool, error)
	CountNumberOfLikes(linkID int) (int, error)
}

type repository struct {
	db *sql.DB
}

// NewRepo -
func NewRepo(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r repository) Create(like *entities.Like) error {
	_, err := r.db.Exec(`INSERT INTO "like" VALUES($1, $2, $3)`,
		like.LinkID,
		like.LikerID,
		like.CreatedAt,
	)

	return err
}
func (r repository) Read(linkID int) ([]*entities.Like, error) {
	rows, err := r.db.Query(`SELECT * FROM "like" WHERE link_id=$1`, linkID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	likes := make([]*entities.Like, 0)

	for rows.Next() {
		var like entities.Like
		err = rows.Scan(&like.LinkID, &like.LikerID, &like.CreatedAt)
		if err != nil {
			return nil, err
		}
		likes = append(likes, &like)
	}

	return likes, err
}

func (r repository) CountNumberOfLikes(linkID int) (int, error) {
	var n int
	err := r.db.QueryRow(`SELECT COUNT(link_id) FROM "like" WHERE link_id=$1`, linkID).Scan(&n)
	if err != nil {
		return -1, err
	}

	return n, nil
}

func (r repository) Delete(linkID, likerID int) (bool, error) {
	res, err := r.db.Exec(`DELETE FROM "like" WHERE link_id=$1 AND liker_id=$2`, linkID, likerID)
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
