package routes

import (
	"start/cmd/handlers"

	"github.com/labstack/echo/v4"
)

func BotApiRoutes(group *echo.Group) {

	group.POST("/commands", handlers.BotApiCommands)
	group.GET("/messages/:id", handlers.GetMessageById)

}
