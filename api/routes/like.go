package routes

import (
	"strconv"

	"github.com/bagus2x/fainmi-be/api/middleware"
	"github.com/bagus2x/fainmi-be/pkg/like"
	"github.com/gofiber/fiber/v2"
)

// Like -
func Like(app fiber.Router, service like.Service, auth middleware.Authentication) {
	v1 := app.Group("/api/v1/like")

	v1.Post("/:link_id", auth.Auth, createLike(service))
	v1.Get("/:link_id", getLikes(service))
	v1.Get("/:link_id/total", getNumberOfLikes(service))
	v1.Delete("/:link_id", auth.Auth, deleteLike(service))
}

func createLike(service like.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		linkID, err := strconv.Atoi(c.Params("link_id"))
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		likerID := c.Locals("profile_id").(int)

		err = service.AddLike(linkID, likerID)
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

func getLikes(service like.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		linkID, err := strconv.Atoi(c.Params("link_id"))
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		res, err := service.GetLikes(linkID)
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

func getNumberOfLikes(service like.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		linkID, err := strconv.Atoi(c.Params("link_id"))
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		res, err := service.GetNumberOfLikes(linkID)
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

func deleteLike(service like.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		linkID, err := strconv.Atoi(c.Params("link_id"))
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		likerID := c.Locals("profile_id").(int)

		err = service.DeleteLike(linkID, likerID)
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
