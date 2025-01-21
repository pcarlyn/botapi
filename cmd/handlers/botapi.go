package handlers

import (
	"fmt"
	"start/internal/models"
	"start/internal/utils"
	"start/internal/utils/requests/answers"
	"start/internal/utils/requests/states"
	"strings"

	"github.com/labstack/echo/v4"
)

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

	answers, _ := answers.GetAnswers("cmd", cmd)
	state, _ := states.GetStatesById(tgid)
	// variables, _ := variables.GetVariables(tgid)

	for _, ans := range answers {
		if state.State == ans.State {
			answer := ans
			utils.BuildAnswer(&answer, &resp)
		}
	}

	return c.JSON(200, resp)
}

func GetMessageById(c echo.Context) error {

	id := c.QueryParam("id")

	return c.JSON(200, "GetMessageById"+id)
}
