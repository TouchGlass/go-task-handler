package handlers

import (
	"BDproj/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	service service.TaskService
}

// создание хэндлеров
func NewHandler(s service.TaskService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) PostTask(c echo.Context) error {
	var task service.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not bind request body"})
	}

	err := h.service.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create task"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Task created successfully"})

}

func (h Handler) GetTasks(c echo.Context) error {
	tasks, err := h.service.GetTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h Handler) PatchTask(c echo.Context) error {
	var task service.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not bind request body"})
	}

	id := c.Param("id")

	if err := h.service.PatchTask(id, task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update task"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Task updated successfully"})

}

func (h Handler) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	err := h.service.DeleteTaskByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Task deleted successfully"})
}
