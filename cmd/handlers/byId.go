package handlers

import (
	"start/internal/models"
	"start/internal/utils"
	"start/internal/utils/requests/answers"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetMessageById godoc
// @Summary Получить сообщение по ID
// @Description Возвращает объект сообщения по его идентификатору
// @Tags BotAPI
// @Accept  json
// @Produce  json
// @Param id path int true "ID сообщения"
// @Success 200 {object} models.ResponseAnswer
// @Failure 400 {object} map[string]interface{} "Ошибка при обработке запроса"
// @Failure 404 {object} map[string]interface{} "Сообщение не найдено"
// @Router /botapi/v1/messages/{id} [get]
func GetMessageById(c echo.Context) error {

	strId := c.Param("id")
	id, _ := strconv.Atoi(strId)

	answer, statusCode := answers.GetAnswerById(id)

	var responseAnswer models.ResponseAnswer

	utils.BuildAnswerV2(&answer, &responseAnswer)

	return c.JSON(statusCode, responseAnswer)
}
