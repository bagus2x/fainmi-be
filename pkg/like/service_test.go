package like

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getService() Service {
	return NewService(NewRepo(db()))
}

func TestServiceAddLike(t *testing.T) {
	s := getService()
	err := s.AddLike(1, 5)
	assert.NoError(t, err)
}

func TestServiceDeleteLike(t *testing.T) {
	s := getService()
	err := s.DeleteLike(1, 2)
	assert.NoError(t, err)
}

func TestServiceGetLike(t *testing.T) {
	s := getService()
	like, err := s.GetLikes(1)
	assert.NoError(t, err)
	assert.NotNil(t, like)
}

func TestServiceGetNumberOfLikes(t *testing.T) {
	service := getService()
	n, err := service.GetNumberOfLikes(1)
	assert.NoError(t, err)
	t.Log(n)
	assert.Positive(t, n)
}
