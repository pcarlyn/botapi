package handlers

import (
	"start/internal/models"
	"start/internal/utils"
	"start/internal/utils/requests/answers"
	"start/internal/utils/requests/states"
	"start/internal/utils/requests/variables"

	"github.com/labstack/echo/v4"
)

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
