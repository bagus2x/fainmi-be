package routes

import (
	"strconv"

	"github.com/bagus2x/fainmi-be/pkg/button"
	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// Button -
func Button(app fiber.Router, service button.Service) {
	v1 := app.Group("/api/v1/button")

	v1.Post("/", createButton(service))
	v1.Get("/", getButtons(service))
	v1.Get("/:btn_id", getButton(service))
	v1.Put("/:btn_id", updateButton(service))
	v1.Delete("/btn_id", deleteButton(service))
}

func createButton(service button.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.CreateButtonRequest
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		err = service.AddButton(&req)
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

func getButton(service button.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		btnID, err := strconv.Atoi(c.Params("btn_id"))
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		res, err := service.GetButton(btnID)
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

func getButtons(service button.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res, err := service.GetButtons()
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

func updateButton(service button.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.UpdateButtonRequest
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		btnID, err := strconv.Atoi(c.Params("btn_id"))
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		err = service.UpdateButton(btnID, &req)
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

func deleteButton(service button.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		btnID, err := strconv.Atoi(c.Params("btn_id"))
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		err = service.DeleteButton(btnID)
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
