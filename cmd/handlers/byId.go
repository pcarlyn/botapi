package handlers

import (
	"start/internal/models"
	"start/internal/utils"
	"start/internal/utils/requests/answers"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetMessageById(c echo.Context) error {

	strId := c.Param("id")
	id, _ := strconv.Atoi(strId)

	answer, statusCode := answers.GetAnswerById(id)

	var responseAnswer models.ResponseAnswer

	utils.BuildAnswerV2(&answer, &responseAnswer)

	return c.JSON(statusCode, responseAnswer)
}
