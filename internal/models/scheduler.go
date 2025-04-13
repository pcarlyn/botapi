package models

import (
	"time"
)

type SchedulerMsg struct {
	Content   string    `json:"content"`
	SendAt    time.Time `json:"send_at"`
	Channel   string    `json:"channel"`
	Recipient string    `json:"recipient"`
	Tag       string    `json:"tag"`
}

type SchedulerAnswer struct {
	Content   string    `json:"content"`
	SendAt    time.Time `json:"send_at"`
	Channel   string    `json:"channel"`
	Recipient string    `json:"recipient"`
	Tag       string    `json:"tag"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateMsg struct {
	Msg string `json:"msg"`
}
