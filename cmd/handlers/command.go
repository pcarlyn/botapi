package handlers

import (
	"fmt"
	"start/internal/models"
	"start/internal/utils"
	"start/internal/utils/requests/answers"
	"start/internal/utils/requests/states"
	"start/internal/utils/requests/variables"
	"strings"

	"github.com/labstack/echo/v4"
)

// BotApiCommands godoc
// @Summary Обработка команд от Telegram-бота
// @Description Получает команду от пользователя (например, /start), фильтрует ответ и возвращает подходящее сообщение, обновляя состояние и переменные
// @Tags BotAPI
// @Accept  json
// @Produce  json
// @Param message body models.Result true "Входящее сообщение от Telegram"
// @Success 200 {object} models.ControllerResponce
// @Failure 400 {object} map[string]interface{} "Ошибка при обработке запроса"
// @Router /botapi/v1/commands [post]
func BotApiCommands(c echo.Context) error {

	var userData models.Result
	var resp models.ControllerResponce
	fmt.Println(c.Request().Body)
	if err := c.Bind(&userData); err != nil {
		return err
	}

	tgid := userData.Message.From.ID

	cmdMsg := userData.Message.Text

	cmd := strings.Replace(cmdMsg, "/", "", 1)
	if cmd == "start" {
		utils.SetState(tgid, "default")
	}

	answers, _ := answers.GetAnswers("cmd", cmd)
	state, _ := states.GetStatesById(tgid)
	variables, _ := variables.GetVariables(tgid)

	answer := utils.FilterAnswers(answers, state, variables)
	utils.BuildAnswer(&answer, &resp)

	utils.SetState(tgid, answer.NextState)
	utils.SetVar(tgid, answer.SetVariable, answer.SetValue)

	return c.JSON(200, resp)
}
