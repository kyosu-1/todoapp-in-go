package main

import (
	"todoapp-in-go/config"
	"todoapp-in-go/repository/gorm"
	"todoapp-in-go/usecase"
	"todoapp-in-go/router"

	"github.com/labstack/echo/v4"
)

func main() {
	taskRepository := gorm.NewTaskRepository(config.NewDB())
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := router.NewTaskHandler(taskUsecase)

	e := echo.New()
	router.InitRouting(e, taskHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
