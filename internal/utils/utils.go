package utils

import (
	"fmt"
	"net/http"
	"start/internal/models"
	"start/internal/utils/requests/states"
	"start/internal/utils/requests/variables"
	"strconv"
)

func BuildAnswer(req *models.ResponseAnswer, resp *models.ControllerResponce) {
	resp.Answer = req.Answer
	resp.Delay = req.Delay
	resp.Keyboard = req.Keyboard
	resp.IsKb = req.IsKb
	resp.IsNextMsg = req.IsNextMessage
	if req.IsNextMessage {
		resp.Id = req.Id
	} else {
		resp.Id = 0
	}
	fmt.Println(resp.IsNextMsg)
	fmt.Printf("%+v\n", resp)
}

func BuildAnswerV2(req *models.GetResponseAnswer, resp *models.ResponseAnswer) {
	resp.Answer = req.Data.Answer
	resp.Delay = req.Data.Delay
	resp.Keyboard = req.Data.Keyboard
	resp.IsKb = req.Data.IsKb
	resp.IsNextMessage = req.Data.IsNextMessage
	resp.NextMessage = req.Data.NextMessage
	if req.Data.IsNextMessage {
		resp.Id = req.Data.Id
	} else {
		resp.Id = 0
	}
	fmt.Println(resp.IsNextMessage)
	fmt.Printf("%+v\n", resp)
}

func FilterAnswers(answers []models.ResponseAnswer, state models.RState, variables []models.VariablesResponse) models.ResponseAnswer {

	var answer models.ResponseAnswer
	messageCandidates := make(map[int]bool)
	filtredAnswers := FilterByStates(&messageCandidates, answers, state)
	answer = FilterByConditions(filtredAnswers, variables, &messageCandidates)
	fmt.Println(filtredAnswers)

	return answer
}

func WriteAdminMessage([]int) {
	fmt.Println("WriteAdminMessage")
}

func FilterByStates(messageCandidates *map[int]bool, answers []models.ResponseAnswer, state models.RState) []models.ResponseAnswer {

	var resultAnswers []models.ResponseAnswer

	for _, answer := range answers {
		if answer.State == "" {
			(*messageCandidates)[answer.Id] = true
			resultAnswers = append(resultAnswers, answer)
			continue
		}
		if answer.State == state.State {
			(*messageCandidates)[answer.Id] = true
			resultAnswers = append(resultAnswers, answer)
		}
	}

	return resultAnswers
}

func FilterByConditions(answers []models.ResponseAnswer, variables []models.VariablesResponse, messageCandidates *map[int]bool) models.ResponseAnswer {
	result := models.ResponseAnswer{}

	for _, answer := range answers {
		for _, condition := range answer.Conditions {
			FilterMessageByCondition(condition, variables, messageCandidates, answer.Id)
		}
	}

	fmt.Println(messageCandidates)
	fmt.Printf("\n\n\n")

	count := 0
	MessageToSend := make([]int, 0, 1)

	for id, value := range *messageCandidates {
		if value {
			count++
			MessageToSend = append(MessageToSend, id)
			result = GetMessageById(answers, id)
		}
	}
	if count > 1 {
		WriteAdminMessage(MessageToSend)
	}

	return result
}

func GetMessageById(answers []models.ResponseAnswer, id int) models.ResponseAnswer {
	for _, answer := range answers {
		if answer.Id == id {
			return answer
		}
	}
	return models.ResponseAnswer{}
}

func FilterMessageByCondition(condition models.Condition, variables []models.VariablesResponse, messageCandidates *map[int]bool, message_id int) {

	flag := false
	for _, variable := range variables {
		if variable.Name == condition.Variable {
			flag = true
		}
	}

	if !flag && (condition.Operation == ">" || condition.Operation == "<" || condition.Operation == "=") {
		(*messageCandidates)[message_id] = false
		return
	}

	for _, variable := range variables {
		if variable.Name == condition.Variable {
			switch condition.Operation {
			case "=":
				if variable.Value != condition.Value {
					(*messageCandidates)[message_id] = false
				}
			case ">":
				intvar, _ := strconv.Atoi(variable.Value)
				intcond, _ := strconv.Atoi(condition.Value)
				if intvar <= intcond {
					(*messageCandidates)[message_id] = false
				}
			case "<":
				intvar, _ := strconv.Atoi(variable.Value)
				intcond, _ := strconv.Atoi(condition.Value)
				if intvar >= intcond {
					(*messageCandidates)[message_id] = false
				}
			case "!=":
				if variable.Value == condition.Value {
					(*messageCandidates)[message_id] = false
				}
			}
		}
	}
}

func SetVar(id int, name string, value string) {
	if name == "" {
		return
	}

	newVariable := models.Variable{
		ProfileId: id,
		Name:      name,
		Value:     value,
	}
	vars, _ := variables.GetVariables(id)
	flag := false
	for _, variable := range vars {
		if variable.Name == name {
			flag = true
		}
	}
	if !flag {
		variables.PostVariables(newVariable)
	} else {
		variables.PatchVariables(newVariable)
	}
}

func SetState(id int, state string) {
	if state == "" {
		return
	}
	newState := models.RState{
		ProfileId: id,
		State:     state,
	}
	_, status := states.GetStatesById(id)
	if status == http.StatusNotFound {
		states.PostStates(newState)
	} else {
		states.PatchStates(newState)
	}
}
