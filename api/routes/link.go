package routes

import (
	"strconv"

	"github.com/bagus2x/fainmi-be/api/middleware"
	"github.com/bagus2x/fainmi-be/pkg/link"
	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// Link -
func Link(app fiber.Router, service link.Service, auth middleware.Authentication) {
	v1 := app.Group("/api/v1/link")

	v1.Post("/", auth.Auth, createLink(service))
	v1.Get("/", auth.Auth, getLinks(service))
	v1.Get("/:link_id", auth.Auth, getLink(service))
	v1.Put("/:link_id", auth.Auth, updateLink(service))
	v1.Get("/public/:username", getPublicLinks(service))
	v1.Put("/", auth.Auth, updateLinksOrder(service))
	v1.Delete("/:link_id", auth.Auth, deleteLink(service))
}

func createLink(service link.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.CreateLinkRequest
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		profileID := c.Locals("profile_id").(int)

		res, err := service.CreateLink(profileID, &req)
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

func getLink(service link.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		linkID, err := strconv.Atoi(c.Params("link_id"))
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		profileID := c.Locals("profile_id").(int)

		res, err := service.GetLink(linkID, profileID)
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

func getPublicLinks(service link.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Params("username")

		res, err := service.GetPublicLinks(username)
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

func getLinks(service link.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		profileID := c.Locals("profile_id").(int)

		res, err := service.GetLinks(profileID)
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

func updateLink(service link.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.LinkUpdateReq
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		linkID, err := strconv.Atoi(c.Params("link_id"))
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		profileID := c.Locals("profile_id").(int)

		err = service.UpdateLink(linkID, profileID, &req)
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

func updateLinksOrder(service link.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.LinksOrder
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		profileID := c.Locals("profile_id").(int)

		err = service.UpdateLinksOrder(profileID, req)
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

func deleteLink(service link.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		linkID, err := strconv.Atoi(c.Params("link_id"))
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		profileID := c.Locals("profile_id").(int)

		err = service.DeleteLink(linkID, profileID)
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
