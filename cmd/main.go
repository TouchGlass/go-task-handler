package main

import (
	"BDproj/internal/db"
	"BDproj/internal/handlers"
	"BDproj/internal/taskService"
	userService2 "BDproj/internal/userService"
	"BDproj/internal/web/tasks"
	"BDproj/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	todoRepo := taskService.NewTaskRepository(database)
	todoService := taskService.NewTaskService(todoRepo)
	todoHandler := handlers.NewTaskHandler(todoService)

	uRepo := userService2.NewUserRepository(database)
	uService := userService2.NewUserService(uRepo)
	uHandler := handlers.NewUserHandler(uService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictTaskHandler := tasks.NewStrictHandler(todoHandler, nil)
	tasks.RegisterHandlers(e, strictTaskHandler)

	strictUserHandler := users.NewStrictHandler(uHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
