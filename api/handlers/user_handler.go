package handlers

import (
	"log"
	"net/http"
	"timesheet-manager-backend/api/presenter"
	"timesheet-manager-backend/pkg/entities"
	"timesheet-manager-backend/pkg/user"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// AddUser is handler/controller which creates Users
func AddUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.User
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		if requestBody.Email == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(errors.New(
				"Please specify user details")))
		}
		log.Printf(requestBody.Email)
		log.Printf(requestBody.Password)
		result, err := service.InsertUser(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(result))
	}
}

// UpdateUser is handler/controller which updates data of Users
func UpdateUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.User
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		result, err := service.UpdateUser(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(result))
	}
}

// RemoveUser is handler/controller which removes users
func RemoveUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		userId := requestBody.ID
		err = service.RemoveUser(userId)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "Updated Successfully!",
			"err":    nil,
		})
	}
}

// LoginUser is handler/controller which logins user
func LoginUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody presenter.LoginRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		var result *entities.User
		result, err = service.LoginUser(requestBody.Email, requestBody.Password)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(result))
	}
}

// GetUsers is handler/controller which lists all Users
func GetUsers(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchUsers()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UsersSuccessResponse(fetched))
	}
}
