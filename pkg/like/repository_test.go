package like

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

var f = entities.Like{
	LinkID:    10,
	LikerID:   2,
	CreatedAt: time.Now().Unix(),
}

func TestRepoCreate(t *testing.T) {
	repo := NewRepo(db())
	err := repo.Create(&f)
	assert.NoError(t, err)
}

func TestRepoRead(t *testing.T) {
	repo := NewRepo(db())
	bg, err := repo.Read(2)
	t.Log(bg)
	t.Log(err)
	assert.NotNil(t, bg)
	assert.NoError(t, err)
}

func TestRepoDelete(t *testing.T) {
	repo := NewRepo(db())
	deleted, err := repo.Delete(1, 2)
	t.Log(err)
	assert.Condition(t, func() (success bool) {
		return deleted
	})
	assert.NoError(t, err)
}

func TestRepoCountNumberOfLikes(t *testing.T) {
	repo := NewRepo(db())
	n, err := repo.CountNumberOfLikes(1)
	assert.NoError(t, err)
	t.Log(n)
	assert.Positive(t, n)
}
