package main

import (
	"BDproj/internal/db"
	"BDproj/internal/handlers"
	"BDproj/internal/service"
	"BDproj/internal/web/tasks"

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
	handler := handlers.NewHandler(todoService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
