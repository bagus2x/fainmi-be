package background

import (
	"testing"

	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/stretchr/testify/assert"
)

func getService() Service {
	return NewService(NewRepo(db()))
}

func TestServiceAddBackground(t *testing.T) {
	s := getService()
	err := s.AddBackground(&models.CreateBackgroundReq{
		Name:        "ijo",
		Description: "ijo",
	})

	assert.NoError(t, err)
}

func TestServiceGetBackground(t *testing.T) {
	s := getService()
	bg, err := s.GetBackground(1)
	assert.NoError(t, err)
	assert.NotNil(t, bg)
}

func TestServiceUpdateBackground(t *testing.T) {
	s := getService()
	err := s.UpdateBackground(1, &models.UpdateBackgroundReq{
		Name:        "Dark",
		Description: "Kegelapan dunia",
	})
	assert.NoError(t, err)
}

func TestServiceDeleteBackground(t *testing.T) {
	s := getService()
	err := s.DeleteBackground(2)
	assert.NoError(t, err)
}
