package style

import (
	"testing"

	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/stretchr/testify/assert"
)

func getService() Service {
	return NewService(NewRepo(db()))
}

func TestServiceCreateStyle(t *testing.T) {
	s := getService()
	err := s.CreateStyle(1, &models.StyleReq{
		BackgroundID: 0,
		ButtonID:     0,
		FontID:       0,
	})
	t.Log(err)
	assert.NoError(t, err)
}

func TestServiceGetStyle(t *testing.T) {
	s := getService()
	style, err := s.GetStyle(1)
	assert.NoError(t, err)
	assert.NotNil(t, style)
	t.Log(style)
}

func TestServiceGetStyleDetail(t *testing.T) {
	s := getService()
	styleDetail, err := s.GetStyleDetail(1)
	assert.NoError(t, err)
	assert.NotNil(t, styleDetail)
	t.Log(styleDetail)
}

func TestServiceUpdateStyle(t *testing.T) {
	s := getService()
	err := s.UpdateStyle(1, &models.StyleReq{
		BackgroundID: 1,
		ButtonID:     0,
		FontID:       0,
	})
	t.Log(err)
	assert.Error(t, err)
}

func TestServiceDeleteStyle(t *testing.T) {
	s := getService()
	err := s.DeleteStyle(1)
	assert.Nil(t, err)
}
