package middleware

import (
	"strings"

	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/bagus2x/fainmi-be/pkg/profile"
	"github.com/gofiber/fiber/v2"
)

type r models.Response

// Authentication middleware
type Authentication interface {
	Auth(c *fiber.Ctx) error
}

type auth struct {
	profileService profile.Service
}

// NewAuth -
func NewAuth(profile profile.Service) Authentication {
	return &auth{profileService: profile}
}

func (a auth) Auth(c *fiber.Ctx) error {
	// Format: [Bearer <token>]
	authorization := strings.Split(c.Get("Authorization"), " ")
	if len(authorization) < 2 {
		return c.Status(401).JSON(r{
			Message: "Authorization header must be filled",
		})
	}
	accessToken := authorization[1]
	claims, err := a.profileService.ParseAccessToken(accessToken)
	if err != nil {
		return c.Status(401).JSON(r{
			Message: err.Error(),
		})
	}

	c.Locals("profile_id", claims.ProfileID)
	return c.Next()
}
