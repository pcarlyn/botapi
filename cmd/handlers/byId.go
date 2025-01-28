package handlers

import (
	"start/internal/utils/requests/answers"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetMessageById(c echo.Context) error {

	strId := c.Param("id")
	id, _ := strconv.Atoi(strId)

	answer, statusCode := answers.GetAnswerById(id)

	return c.JSON(statusCode, answer)
}
