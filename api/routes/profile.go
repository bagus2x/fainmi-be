package routes

import (
	"fmt"
	"path"
	"path/filepath"
	"time"

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
	v1.Get("/", auth.Auth, getProfile(service))
	v1.Put("/update", auth.Auth, updateProfile(service))
	v1.Patch("/photo", auth.Auth, updatePhotoProfile(service))
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

func getProfile(service profile.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		profileID := c.Locals("profile_id").(int)

		res, err := service.GetProfile(profileID)
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

func updatePhotoProfile(service profile.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("photo")
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		profileID := c.Locals("profile_id").(int)

		fileName := fmt.Sprintf("%d-%d%s", profileID, time.Now().Unix(), filepath.Ext(file.Filename))
		dest := path.Join("public", "photo", fileName)
		err = c.SaveFile(file, fmt.Sprintf("./%s", dest))
		if err != nil {
			return c.Status(code(err)).JSON(r{
				Message: err.Error(),
			})
		}

		service.UpdatePhoto(profileID, dest)

		return c.JSON(r{
			Success: true,
			Data:    dest,
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
