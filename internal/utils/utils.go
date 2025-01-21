package utils

import (
	"fmt"
	"start/internal/models"
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
