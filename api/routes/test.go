package routes

import (
	"github.com/bagus2x/fainmi-be/api/middleware"
	"github.com/gofiber/fiber/v2"
)

// Test -
func Test(app fiber.Router, auth middleware.Authentication) {
	app.Get("/test", auth.Auth, test)
}

func test(c *fiber.Ctx) error {
	return c.JSON(c.Locals("profile_id").(int))
}

func m1(c *fiber.Ctx) error {
	c.Locals("nama", "bagus")
	return c.Next()
}
