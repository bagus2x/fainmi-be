package routes

import (
	"strconv"

	"github.com/bagus2x/fainmi-be/pkg/background"
	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// Background -
func Background(app fiber.Router, service background.Service) {
	v1 := app.Group("/api/v1/background")

	v1.Post("/", createBackground(service))
	v1.Get("/", getBackgrounds(service))
	v1.Get("/:bg_id", getBackground(service))
	v1.Put("/:bg_id", updateBackground(service))
	v1.Delete("/bg_id", deleteBackground(service))
}

func createBackground(service background.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.CreateBackgroundRequest
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		err = service.AddBackground(&req)
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

func getBackground(service background.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bID := c.Params("bg_id")
		bgID, err := strconv.Atoi(bID)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		res, err := service.GetBackground(bgID)
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

func getBackgrounds(service background.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res, err := service.GetBackgrounds()
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

func updateBackground(service background.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.UpdateBackgroundRequest
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		bID := c.Params("bg_id")
		bgID, err := strconv.Atoi(bID)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		err = service.UpdateBackground(bgID, &req)
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

func deleteBackground(service background.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bID := c.Params("bg_id")
		bgID, err := strconv.Atoi(bID)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		err = service.DeleteBackground(bgID)
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
