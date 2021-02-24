package link

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/bagus2x/fainmi-be/pkg/entities"
	"github.com/bagus2x/fainmi-be/pkg/models/errors"
)

// Repository link
type Repository interface {
	Create(link *entities.Link) error
	Read(linkID, profileID int) (*entities.Link, error)
	ReadByProfileID(profileID int) ([]*entities.Link, error)
	ReadByUsername(username string) ([]*entities.Link, error)
	Update(linkID, profileID int, link *entities.Link) (bool, error)
	UpdateDisplay(linkID, profileID int, value bool) (bool, error)
	Delete(linkID, profileID int) (bool, error)
	UpdateOrder(profileID int, orders []*entities.Order) error
}

type repository struct {
	db *sql.DB
}

// NewRepo -
func NewRepo(db *sql.DB) Repository {
	return &repository{db: db}
}

// Create method will assign LinkID
func (r repository) Create(link *entities.Link) error {
	err := r.db.QueryRow(`INSERT INTO link VALUES(DEFAULT, $1, $2, $3, $4, $5, $6, $7) RETURNING link_id`,
		link.ProfileID,
		link.Order,
		link.Title,
		link.URL,
		link.Display,
		link.CreatedAt,
		link.UpdatedAt,
	).Scan(&link.LinkID)

	return err
}
func (r repository) Read(linkID, profileID int) (*entities.Link, error) {
	var link entities.Link
	q := `SELECT * FROM link WHERE link_id=$1 AND profile_id=$2`
	err := r.db.QueryRow(q, linkID, profileID).Scan(
		&link.LinkID,
		&link.ProfileID,
		&link.Order,
		&link.Title,
		&link.URL,
		&link.Display,
		&link.CreatedAt,
		&link.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &link, err
}

func (r repository) ReadByProfileID(profileID int) ([]*entities.Link, error) {
	q := `SELECT link_id, profile_id, "order", title, url, display FROM link WHERE profile_id=$1 ORDER BY "order", updated_at DESC`
	rows, err := r.db.Query(q, profileID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	links := make([]*entities.Link, 0)

	for rows.Next() {
		var link entities.Link
		err = rows.Scan(&link.LinkID, &link.ProfileID, &link.Order, &link.Title, &link.URL, &link.Display)
		if err != nil {
			return nil, err
		}
		links = append(links, &link)
	}

	return links, nil
}

func (r repository) ReadByUsername(username string) ([]*entities.Link, error) {
	q := ` SELECT link_id, profile_id, "order", title, url, display FROM link WHERE profile_id=
		   (SELECT profile_id FROM profile WHERE username=$1) AND display=true`
	rows, err := r.db.Query(q, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	links := make([]*entities.Link, 0)

	for rows.Next() {
		var link entities.Link
		err = rows.Scan(&link.LinkID, &link.ProfileID, &link.Order, &link.Title, &link.URL, &link.Display)
		if err != nil {
			return nil, err
		}
		links = append(links, &link)
	}

	return links, nil
}

// Update method will update order, title, url, display, updated_at
func (r repository) Update(linkID, profileID int, link *entities.Link) (bool, error) {
	q := `UPDATE link SET "order"=$1, title=$2, url=$3, display=$4, updated_at=$5 WHERE link_id=$6 AND profile_id=$7`
	res, err := r.db.Exec(q, link.Order, link.Title, link.URL, link.Display, link.UpdatedAt, linkID, profileID)
	if err != nil {
		return false, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (r repository) UpdateDisplay(linkID, profileID int, value bool) (bool, error) {
	res, err := r.db.Exec(`UPDATE link SET display=$1 WHERE link_id=$2 AND profile_id=$3`, value, linkID, profileID)
	if err != nil {
		return false, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (r repository) Delete(linkID, profileID int) (bool, error) {
	res, err := r.db.Exec(`DELETE FROM link WHERE link_id=$1 AND profile_id=$2`, linkID, profileID)
	if err != nil {
		return false, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (r repository) UpdateOrder(profileID int, orders []*entities.Order) error {
	q := `UPDATE link AS l SET "order"=nl.order FROM(values `
	for i, order := range orders {
		q += `(` + strconv.Itoa(order.LinkID) + `,` + strconv.Itoa(profileID) + `,` + strconv.Itoa(order.Order) + `)`
		if len(orders)-1 != i {
			q += `,`
		}
	}
	q += `)AS nl(link_id, profile_id, "order") WHERE l.link_id=nl.link_id AND l.profile_id=nl.profile_id`

	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	res, err := tx.Exec(q)
	if err != nil {
		tx.Rollback()
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if rowsAffected != int64(len(orders)) {
		tx.Rollback()
		return errors.ErrorMessage(errors.ErrBadRequest, "Link ID not Found")
	}
	err = tx.Commit()

	return nil
}
