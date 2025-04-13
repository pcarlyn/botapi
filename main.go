package main

import (
	"fmt"
	"start/cmd/routes"

	_ "start/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           START-API
// @version         1.0
// @description     API Server for Frontend and Bot
// @BasePath  /
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.EchoWrapHandler(func(c *echoSwagger.Config) {
		c.URLs = []string{fmt.Sprintf("http://%s:%s/swagger/doc.json", "localhost", "8080")}
	}))

	botapi := e.Group("/botapi/v1")

	frontapi := e.Group("/frontapi/v1")

	routes.BotApiRoutes(botapi)

	routes.SchedulerApiRoutes(frontapi)

	routes.ProcessrApiRoutes(frontapi)

	e.Logger.Fatal(e.Start(":8080"))

}
