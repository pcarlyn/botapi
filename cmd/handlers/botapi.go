package handlers

import "github.com/labstack/echo/v4"

func BotApiCommands(c echo.Context) error {

	return c.JSON(200, "BotApiCommands")
}

func GetMessageById(c echo.Context) error {

	return c.JSON(200, "GetMessageById")
}
