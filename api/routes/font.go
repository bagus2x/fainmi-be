package routes

import (
	"strconv"

	"github.com/bagus2x/fainmi-be/pkg/font"
	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// Font -
func Font(app fiber.Router, service font.Service) {
	v1 := app.Group("/api/v1/font")

	v1.Post("/", createFont(service))
	v1.Get("/:font_id", getFont(service))
	v1.Get("/", getFonts(service))
	v1.Put("/:font_id", updateFont(service))
	v1.Delete("/:font_id", deleteFont(service))
}

func createFont(service font.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.CreateFontRequest
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		err = service.AddFont(&req)
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

func getFont(service font.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fontID, err := strconv.Atoi(c.Params("font_id"))
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		res, err := service.GetFont(fontID)
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

func getFonts(service font.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res, err := service.GetFonts()
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

func updateFont(service font.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.UpdateFontRequest
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		fontID, err := strconv.Atoi(c.Params("font_id"))
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		err = service.UpdateFont(fontID, &req)
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

func deleteFont(service font.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fontID, err := strconv.Atoi(c.Params("font_id"))
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		err = service.DeleteFont(fontID)
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
