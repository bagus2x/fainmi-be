package profile

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

var p = &entities.Profile{
	Username:  "mukhlis",
	Email:     "mukhlis@gmail.com",
	Password:  "mukhlis123",
	CreatedAt: time.Now().Unix(),
	UpdatedAt: time.Now().Unix(),
}

func TestRepoCreate(t *testing.T) {
	repo := NewRepo(db())
	err := repo.Create(p)
	t.Log(p.ProfileID)
	assert.NotEqual(t, 0, p.ProfileID)
	assert.NoError(t, err)
}

func TestRepoRead(t *testing.T) {
	repo := NewRepo(db())
	profile, err := repo.Read(0)
	assert.Nil(t, err)
	assert.NotNil(t, profile)
	t.Logf("%+v", profile)
}

func TestRepoUpdate(t *testing.T) {
	repo := NewRepo(db())
	updated, err := repo.Update(1, p)
	assert.Condition(t, func() (success bool) {
		return updated
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
