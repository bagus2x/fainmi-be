package button

import (
	"testing"

	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/stretchr/testify/assert"
)

func getService() Service {
	return NewService(NewRepo(db()))
}

func TestServiceAddButton(t *testing.T) {
	s := getService()
	err := s.AddButton(&models.CreateButtonRequest{
		Name:        "rounded",
		Description: "rounded",
	})

	assert.NoError(t, err)
}

func TestServiceGetButton(t *testing.T) {
	s := getService()
	bg, err := s.GetButton(1)
	assert.NoError(t, err)
	assert.NotNil(t, bg)
}

func TestServiceGetButtons(t *testing.T) {
	s := getService()
	btn, err := s.GetButtons()
	t.Log(len(btn))
	assert.NoError(t, err)
	assert.NotNil(t, btn)
}

func TestServiceUpdateButton(t *testing.T) {
	s := getService()
	err := s.UpdateButton(1, &models.UpdateButtonRequest{
		Name:        "Sharp",
		Description: "Tajam",
	})
	assert.NoError(t, err)
}

func TestServiceDeleteButton(t *testing.T) {
	s := getService()
	err := s.DeleteButton(2)
	assert.NoError(t, err)
}
