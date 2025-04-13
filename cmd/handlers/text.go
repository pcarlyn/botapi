package handlers

import (
	"start/internal/models"
	"start/internal/utils"
	"start/internal/utils/requests/answers"
	"start/internal/utils/requests/states"
	"start/internal/utils/requests/variables"

	"github.com/labstack/echo/v4"
)

// TextHandler godoc
// @Summary Обработка текстовых сообщений от Telegram-бота
// @Description Получает обычное текстовое сообщение от пользователя, подбирает ответ, обновляет состояние и переменные
// @Tags BotAPI
// @Accept  json
// @Produce  json
// @Param message body models.Result true "Текстовое сообщение от Telegram"
// @Success 200 {object} models.ControllerResponce
// @Failure 400 {object} map[string]interface{} "Ошибка при обработке запроса"
// @Router /botapi/v1/messages [post]
func TextHandler(c echo.Context) error {
	var userData models.Result
	var resp models.ControllerResponce

	if err := c.Bind(&userData); err != nil {
		return err
	}

	tgid := userData.Message.From.ID

	txtMsg := userData.Message.Text

	answers, _ := answers.GetAnswers("txt", txtMsg)
	state, _ := states.GetStatesById(tgid)
	variables, _ := variables.GetVariables(tgid)

	answer := utils.FilterAnswers(answers, state, variables)
	utils.BuildAnswer(&answer, &resp)

	utils.SetState(tgid, answer.NextState)
	utils.SetVar(tgid, answer.SetVariable, answer.SetValue)
	return c.JSON(200, resp)
}
