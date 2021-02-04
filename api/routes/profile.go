package routes

import (
	"strconv"

	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/bagus2x/fainmi-be/pkg/profile"
	"github.com/gofiber/fiber/v2"
)

// Profile endpoints
func Profile(app fiber.Router, service profile.Service) {
	v1 := app.Group("/api/v1/profile")

	v1.Post("/signin", signIn(service))
	v1.Post("/signup", signUp(service))
	v1.Put("/update/:id", update(service))
	v1.Delete("/delete/:id", delete(service))
}

func signIn(service profile.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.SignInReq
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		response, err := service.SignIn(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		return c.JSON(r{
			Success: true,
			Data:    response,
		})
	}
}

func signUp(service profile.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.SignUpReq
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		response, err := service.SignUp(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		return c.JSON(r{
			Success: true,
			Data:    response,
		})
	}
}

func update(service profile.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.ProfileUpdateReq
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		pID := c.Params("id")
		id, err := strconv.Atoi(pID)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		response, err := service.UpdateProfile(id, &req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		return c.JSON(r{
			Success: true,
			Data:    response,
		})
	}
}

func delete(service profile.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		pID := c.Params("id")
		id, err := strconv.Atoi(pID)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		err = service.DeleteProfile(id)
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
