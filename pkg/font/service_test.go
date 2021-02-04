package font

import (
	"testing"

	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/stretchr/testify/assert"
)

func getService() Service {
	return NewService(NewRepo(db()))
}

func TestServiceAddFont(t *testing.T) {
	s := getService()
	err := s.AddFont(&models.CreateFontReq{
		Name:        "miring",
		Description: "mirng ke kanan",
	})

	assert.NoError(t, err)
}

func TestServiceGetFont(t *testing.T) {
	s := getService()
	bg, err := s.GetFont(1)
	assert.NoError(t, err)
	assert.NotNil(t, bg)
}

func TestServiceUpdateFont(t *testing.T) {
	s := getService()
	err := s.UpdateFont(1, &models.UpdateFontReq{
		Name:        "Font gelap",
		Description: "Kegelapan dunia",
	})
	assert.NoError(t, err)
}

func TestServiceDeleteFont(t *testing.T) {
	s := getService()
	err := s.DeleteFont(2)
	assert.NoError(t, err)
}
