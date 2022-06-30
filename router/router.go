package router

import (
	"github.com/labstack/echo/v4"
)

// InitRouting routesの初期化
func InitRouting(e *echo.Echo, taskHandler TaskHandler) {
	e.POST("/tasks", taskHandler.PostTask)
	e.GET("/tasks", taskHandler.GetTasks)
	e.GET("/tasks/:id", taskHandler.GetTask)
	e.PUT("/tasks/:id", taskHandler.PutTask)
	e.DELETE("/tasks/:id", taskHandler.DeleteTask)
}
