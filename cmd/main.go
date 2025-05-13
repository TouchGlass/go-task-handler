package main

import (
	"BDproj/internal/db"
	"BDproj/internal/handlers"
	"BDproj/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

/*
	func CreateTask(c echo.Context) error {
		var task Task
		if err := c.Bind(&task); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not bind request body"})
		}

		if err := DB.Create(&task).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save task"})
		}

		return c.JSON(http.StatusCreated, task)
	}

	func GetTask(c echo.Context) error {
		var tasks []Task

		if err := DB.Find(&tasks).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to find tasks"})
		}

		return c.JSON(http.StatusOK, tasks)
	}

	func UpdateTask(c echo.Context) error {
		var task Task
		if err := c.Bind(&task); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not bind request body"})
		}

		idStr := c.Param("id")
		id64, err := strconv.ParseUint(idStr, 10, 0) // Базовый тип uint64
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID must be a positive integer"})
		}
		id := uint(id64) // Приведение к uint

		var task_in_db Task
		if err := DB.First(&task_in_db, "id = ?", id).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Incorrect ID"})
		}

		task_in_db.WhatIsTheTask = task.WhatIsTheTask
		task_in_db.IsDone = task.IsDone

		if err := DB.Save(&task_in_db).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save task"})
		}

		return c.JSON(http.StatusOK, task_in_db)
	}

	func DeleteTask(c echo.Context) error {
		idStr := c.Param("id")
		id64, err := strconv.ParseUint(idStr, 10, 0) // Базовый тип uint64
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID must be a positive integer"})
		}
		id := uint(id64) // Приведение к uint

		result := DB.Delete(&Task{}, id)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete task from database"})
		}

		// Проверяем, была ли удалена хотя бы одна запись
		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
		}
		return c.JSON(http.StatusNoContent, nil)
	}
*/
func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	todoRepository := service.NewTaskRepository(database)
	todoService := service.NewTaskService(todoRepository)
	todoHandlers := handlers.NewHandler(todoService)

	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("api/tasks", todoHandlers.PostTask)
	e.GET("api/tasks", todoHandlers.GetTasks)
	e.PATCH("api/tasks/:id", todoHandlers.PatchTask)
	e.DELETE("api/tasks/:id", todoHandlers.DeleteTask)

	e.Logger.Fatal(e.Start(":8080"))
}
