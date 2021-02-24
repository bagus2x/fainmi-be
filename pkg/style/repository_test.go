package style

import (
	"database/sql"
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

var p = &entities.Style{
	BackgroundID: sql.NullInt32{Valid: true, Int32: 1},
	ButtonID:     sql.NullInt32{Valid: false, Int32: 0},
	FontID:       sql.NullInt32{Valid: false, Int32: 0},
	CreatedAt:    time.Now().Unix(),
	UpdatedAt:    time.Now().Unix(),
}

func TestRepoCreate(t *testing.T) {
	repo := NewRepo(db())
	err := repo.Create(p)
	assert.NoError(t, err)
}

func TestRepoRead(t *testing.T) {
	repo := NewRepo(db())
	style, err := repo.Read(0)
	assert.NoError(t, err)
	assert.NotNil(t, style)
	t.Logf("%+v", style)
}

func TestRepoUpdate(t *testing.T) {
	repo := NewRepo(db())
	isUpated, err := repo.Update(0, p)
	assert.Condition(t, func() (success bool) {
		return isUpated
	})
	assert.Nil(t, err)
}

func TestRepoDelete(t *testing.T) {
	repo := NewRepo(db())
	deleted, err := repo.Delete(0)
	assert.Condition(t, func() (success bool) {
		return deleted
	})
	assert.Nil(t, err)
}

func TestRepoReadStyleDetail(t *testing.T) {
	repo := NewRepo(db())
	style, err := repo.ReadStyleDetail("bagus")
	assert.NoError(t, err)
	assert.NotNil(t, style)
	t.Logf("%+v", style)
}
