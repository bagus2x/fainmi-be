package profile

import (
	"testing"

	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/stretchr/testify/assert"
)

func getService() Service {
	return NewService(NewRepo(db()), "testkey")
}

func TestServiceSignUp(t *testing.T) {
	s := getService()
	res, err := s.SignUp(&models.SignUpRequest{
		Email:    "jaenab@gmail.com",
		Password: "jaenab123",
		Username: "jaenab",
	})
	t.Log(res.Profile.ProfileID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestServiceSignIn(t *testing.T) {
	s := getService()
	res, err := s.SignIn(&models.SignInRequest{
		UsernameOrEmail: "bagus",
		Password:        "bagus123",
	})
	t.Log(res)
	t.Log(err)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestServiceUpdateProfile(t *testing.T) {
	s := getService()
	res, err := s.UpdateProfile(1, &models.ProfileUpdateRequest{
		Email:    "makmur@gmail.com",
		Password: "makmur123",
		Username: "makmur",
		Photo:    "maman.jpg",
	})

	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestServiceDeleteProfile(t *testing.T) {
	s := getService()
	err := s.DeleteProfile(16)
	assert.NoError(t, err)
}
