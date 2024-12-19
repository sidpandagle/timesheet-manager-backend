package handlers

import (
	"net/http"
	"timesheet-manager-backend/api/presenter"
	"timesheet-manager-backend/pkg/entities"
	"timesheet-manager-backend/pkg/task"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func AddTask(service task.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Task
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.TaskErrorResponse(err))
		}
		if requestBody.Task == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TaskErrorResponse(errors.New(
				"Please specify task details")))
		}
		result, err := service.InsertTask(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TaskErrorResponse(err))
		}
		return c.JSON(presenter.TaskSuccessResponse(result))
	}
}

func UpdateTask(service task.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Task
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.TaskErrorResponse(err))
		}
		result, err := service.UpdateTask(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TaskErrorResponse(err))
		}
		return c.JSON(presenter.TaskSuccessResponse(result))
	}
}

func RemoveTask(service task.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.TaskErrorResponse(err))
		}
		taskID := requestBody.ID
		err = service.RemoveTask(taskID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TaskErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}

func GetTasks(service task.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchTasks()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TaskErrorResponse(err))
		}
		return c.JSON(presenter.TasksSuccessResponse(fetched))
	}
}