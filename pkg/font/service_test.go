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
	err := s.AddFont(&models.CreateFontRequest{
		Name:        "miring",
		Description: "mirng ke kanan",
		FontFamily:  "tulisan",
		Href:        "dnwndwddwdwdwd.dwdnwidw",
	})

	assert.NoError(t, err)
}

func TestServiceGetFont(t *testing.T) {
	s := getService()
	bg, err := s.GetFont(1)
	assert.NoError(t, err)
	assert.NotNil(t, bg)
}

func TestServiceGetFonts(t *testing.T) {
	s := getService()
	fnt, err := s.GetFonts()
	t.Log(len(fnt))
	assert.NoError(t, err)
	assert.NotNil(t, fnt)
}

func TestServiceUpdateFont(t *testing.T) {
	s := getService()
	err := s.UpdateFont(1, &models.UpdateFontRequest{
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
