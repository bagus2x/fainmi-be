package link

import (
	"testing"

	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/stretchr/testify/assert"
)

func getService() Service {
	return NewService(NewRepo(db()))
}

func TestServiceCreateLink(t *testing.T) {
	s := getService()
	res, err := s.CreateLink(1, &models.CreateLinkRequest{
		Title:   "my one",
		URL:     "one.com/jaenab",
		Display: true,
		Order:   4,
	})
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestServiceDeleteLink(t *testing.T) {
	s := getService()
	err := s.DeleteLink(2, 1)
	assert.NoError(t, err)
}

func TestServiceGetLink(t *testing.T) {
	s := getService()
	link, err := s.GetLink(3, 1)
	assert.NoError(t, err)
	assert.NotNil(t, link)
}

func TestServiceGetLinks(t *testing.T) {
	s := getService()
	link, err := s.GetLinks(1)
	t.Log(link)
	assert.NoError(t, err)
	assert.NotNil(t, link)
}

func TestServiceUpdateLink(t *testing.T) {
	s := getService()
	err := s.UpdateLink(1, 1, &models.LinkUpdateReq{
		Order:   1009,
		Display: false,
		Title:   "my fb",
		URL:     "www.facebook.com/uni",
	})
	assert.NoError(t, err)
}

func TestServiceUpdateLinkOrder(t *testing.T) {
	s := getService()
	err := s.UpdateLinksOrder(1, models.LinksOrder{{LinkID: 3, Order: 100}})
	assert.NoError(t, err)
}
