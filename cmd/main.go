package main

import (
	"BDproj/internal/db"
	"BDproj/internal/handlers"
	"BDproj/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

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
