package models

type RState struct {
	ProfileId int    `json:"profile_id"`
	State     string `json:"state"`
}

var BaseUrl = "http://test.dev2death.ru"
