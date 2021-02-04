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
	err := s.AddButton(&models.CreateButtonReq{
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

func TestServiceUpdateButton(t *testing.T) {
	s := getService()
	err := s.UpdateButton(1, &models.UpdateButtonReq{
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
