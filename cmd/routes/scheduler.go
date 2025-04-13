package routes

import (
	"start/cmd/handlers"

	"github.com/labstack/echo/v4"
)

func SchedulerApiRoutes(group *echo.Group) {
	group.POST("/schedulers", handlers.PostScheduler)
	group.PUT("/schedulers/:id", handlers.ChangeScheduler)
	group.GET("/schedulers", handlers.GetSchedulers)
	group.GET("/schedulers/:id", handlers.GetSchedulerById)
	group.DELETE("/schedulers/:id", handlers.DeleteScheduler)
	group.PATCH("/schedulers/:id", handlers.UpdateScheduler)
}
