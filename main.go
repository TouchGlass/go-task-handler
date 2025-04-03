package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

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

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Task
	DB.AutoMigrate(&Task{})

	e := echo.New()
	e.POST("api/tasks", CreateTask)
	e.GET("api/tasks", GetTask)

	e.Logger.Fatal(e.Start(":8080"))

	//router := mux.NewRouter()
	//router.HandleFunc("/api/tasks", CreateTask).Methods("POST")
	//router.HandleFunc("/api/tasks", GetTask).Methods("GET")
	//http.ListenAndServe(":8080", router)
}
