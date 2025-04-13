package handlers

import (
	"start/internal/models"
	"start/internal/utils"
	"start/internal/utils/requests/answers"
	"start/internal/utils/requests/states"
	"start/internal/utils/requests/variables"

	"github.com/labstack/echo/v4"
)

// CallbackHandler godoc
// @Summary Обработка callback-запросов от Telegram
// @Description Обрабатывает callback-запрос, фильтрует ответы и обновляет состояние и переменные пользователя
// @Tags BotAPI
// @Accept  json
// @Produce  json
// @Param callback body models.CallBackData true "Данные callback-запроса"
// @Success 200 {object} models.ControllerResponce
// @Failure 400 {object} map[string]interface{} "Ошибка привязки данных"
// @Router /botapi/v1/callbacks [post]
func CallbackHandler(c echo.Context) error {

	var resp models.ControllerResponce
	var CallbackData models.CallBackData

	if err := c.Bind(&CallbackData); err != nil {
		return err
	}
	tgid := CallbackData.CallbackQuery.From.ID

	answers, _ := answers.GetAnswers("clb", CallbackData.CallbackQuery.Message.Text)
	state, _ := states.GetStatesById(tgid)
	variables, _ := variables.GetVariables(tgid)

	answer := utils.FilterAnswers(answers, state, variables)
	utils.BuildAnswer(&answer, &resp)

	utils.SetState(tgid, answer.NextState)
	utils.SetVar(tgid, answer.SetVariable, answer.SetValue)

	return c.JSON(200, resp)

}
