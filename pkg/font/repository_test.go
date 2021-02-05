package font

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

var f = entities.Font{
	Name:        "lucida",
	Description: sql.NullString{String: "lucida is lucida", Valid: true},
	CreatedAt:   time.Now().Unix(),
	UpdatedAt:   time.Now().Unix(),
}

func TestRepoCreate(t *testing.T) {
	repo := NewRepo(db())
	err := repo.Create(&f)
	assert.NoError(t, err)
}

func TestRepoRead(t *testing.T) {
	repo := NewRepo(db())
	bg, err := repo.Read(1) // not found
	t.Log(bg)
	t.Log(err)
	assert.NotNil(t, bg)
	assert.NoError(t, err)
}

func TestRepoReadAll(t *testing.T) {
	repo := NewRepo(db())
	bg, err := repo.ReadAll()
	t.Log(bg)
	t.Log(err)
	assert.NotNil(t, bg)
	assert.NoError(t, err)
}

func TestRepoUpdate(t *testing.T) {
	repo := NewRepo(db())
	_, err := repo.Update(2, &f)
	t.Log(err)
	assert.NoError(t, err)
}

func TestRepoDelete(t *testing.T) {
	repo := NewRepo(db())
	_, err := repo.Delete(2)
	t.Log(err)
	assert.NoError(t, err)
}
