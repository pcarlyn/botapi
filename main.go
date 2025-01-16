package main

import (
	"start/cmd/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	botapi := e.Group("/botapi/v1")

	routes.BotApiRoutes(botapi)

	e.Logger.Fatal(e.Start(":8080"))
}
