package routes

import (
	"github.com/bagus2x/fainmi-be/api/middleware"
	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/bagus2x/fainmi-be/pkg/style"
	"github.com/gofiber/fiber/v2"
)

// Style -
func Style(app fiber.Router, service style.Service, auth middleware.Authentication) {
	v1 := app.Group("/api/v1/style")

	v1.Post("/", auth.Auth, createStyle(service))
	v1.Get("/", auth.Auth, getStyle(service))
	v1.Get("/:username", getStyleDetail(service))
	v1.Put("/", auth.Auth, updateStyle(service))
	v1.Delete("/", auth.Auth, deleteStyle(service))
}

func createStyle(service style.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.StyleRequest
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		profileID := c.Locals("profile_id").(int)

		err = service.CreateStyle(profileID, &req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		return c.JSON(r{
			Success: true,
		})
	}
}

func getStyle(service style.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		profileID := c.Locals("profile_id").(int)

		res, err := service.GetStyle(profileID)
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

func getStyleDetail(service style.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Params("username")

		res, err := service.GetStyleDetail(username)
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

func updateStyle(service style.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.StyleRequest
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		profileID := c.Locals("profile_id").(int)

		err = service.UpdateStyle(profileID, &req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		return c.JSON(r{
			Success: true,
		})
	}
}

func deleteStyle(service style.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		profileID := c.Locals("profile_id").(int)

		err := service.DeleteStyle(profileID)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		return c.JSON(r{
			Success: true,
		})
	}
}
