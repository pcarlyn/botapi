package routes

import (
	"start/cmd/handlers"

	"github.com/labstack/echo/v4"
)

func ProcessrApiRoutes(group *echo.Group) {
	group.POST("/processes", handlers.StartProcess)
	group.GET("/processes", handlers.GetProcesses)
	group.GET("/processes/:id", handlers.GetProcess)
	group.DELETE("/processes/:id", handlers.StopProcess)
}
