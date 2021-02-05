package link

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/bagus2x/fainmi-be/config"
	"github.com/bagus2x/fainmi-be/pkg/entities"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func db() *sql.DB {
	uri := "postgres://bagus2x:admin123@localhost:5432/fainmi"
	db, err := config.DatabaseConnection(uri)
	if err != nil {
		panic(err)
	}

	return db
}

var f = entities.Link{
	ProfileID: 6,
	Order:     0,
	Display:   true,
	Title:     sql.NullString{Valid: true, String: "dnwndiwd"},
	URL:       sql.NullString{Valid: true, String: "wwww.wdnwdnwd.com/wkdnwd"},
	CreatedAt: time.Now().Unix(),
	UpdatedAt: time.Now().Unix(),
}

func TestRepoCreate(t *testing.T) {
	repo := NewRepo(db())
	err := repo.Create(&f)
	t.Log(f.LinkID)
	assert.NoError(t, err)
}

func TestRepoRead(t *testing.T) {
	repo := NewRepo(db())
	bg, err := repo.Read(1, 2)
	t.Log(bg)
	t.Log(err)
	assert.NotNil(t, bg)
	assert.NoError(t, err)
}

func TestRepoUpdate(t *testing.T) {
	repo := NewRepo(db())
	isUpdated, err := repo.Update(80, 2, &f)
	t.Log(err)
	assert.Condition(t, func() (success bool) {
		return isUpdated
	})
	assert.NoError(t, err)
}

func TestRepoDelete(t *testing.T) {
	repo := NewRepo(db())
	isDeleted, err := repo.Delete(3, 1)
	t.Log(err)
	assert.Condition(t, func() (success bool) {
		return isDeleted
	})
	assert.NoError(t, err)
}

func TestRepoReadByProfileID(t *testing.T) {
	repo := NewRepo(db())
	links, err := repo.ReadByProfileID(1)
	assert.NoError(t, err)
	assert.NotNil(t, links)
	for _, v := range links {
		t.Logf("%+v", v)
	}
}

var orders = []*entities.Order{{LinkID: 8, Order: 1290}, {LinkID: 11, Order: 539}, {LinkID: 12, Order: 130}}

func BenchmarkStringSprint(b *testing.B) {
	var values []string
	for _, order := range orders {
		values = append(values, fmt.Sprintf("(%d,%d)", order.LinkID, order.Order))
	}
	q := `UPDATE link AS l SET "order"=nl.order FROM(values %s) AS nl(link_id, "order") WHERE l.link_id=nl.link_id`
	q = fmt.Sprintf(q, strings.Join(values, ","))
	_ = q
}

func BenchmarkStringConcat(b *testing.B) {
	q := `UPDATE link AS l SET "order"=nl.order FROM(values `
	for i, order := range orders {
		q += `(` + strconv.Itoa(order.LinkID) + "," + strconv.Itoa(order.Order) + `)`
		if len(orders)-1 != i {
			q += `,`
		}
	}
	q += `)AS nl(link_id, "order") WHERE l.link_id=nl.link_id`
}

func TestRepoUpdateOrder(t *testing.T) {
	r := NewRepo(db())
	err := r.UpdateOrder(1, orders)
	assert.NoError(t, err)
}

// “Premature optimization is the root of all evil” - Donald Knuth
