package routes

import (
	"github.com/bagus2x/fainmi-be/api/middleware"
	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/bagus2x/fainmi-be/pkg/profile"
	"github.com/gofiber/fiber/v2"
)

// Profile endpoints
func Profile(app fiber.Router, service profile.Service, auth middleware.Authentication) {
	v1 := app.Group("/api/v1/profile")

	v1.Post("/signin", signIn(service))
	v1.Post("/signup", signUp(service))
	v1.Put("/update", auth.Auth, updateProfile(service))
	v1.Delete("/delete", auth.Auth, deleteProfile(service))
}

func signIn(service profile.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.SignInRequest
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		res, err := service.SignIn(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		return c.JSON(r{
			Success: true,
			Data:    res,
		})
	}
}

func signUp(service profile.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.SignUpRequest
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		res, err := service.SignUp(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		return c.JSON(r{
			Success: true,
			Data:    res,
		})
	}
}

func updateProfile(service profile.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.ProfileUpdateRequest
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		profileID := c.Locals("profile_id").(int)

		res, err := service.UpdateProfile(profileID, &req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		return c.JSON(r{
			Success: true,
			Data:    res,
		})
	}
}

func deleteProfile(service profile.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		profileID := c.Locals("profile_id").(int)

		err := service.DeleteProfile(profileID)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		return c.JSON(r{
			Success: true,
			Message: "Delete user successfully",
		})
	}
}
