package models

type ControllerResponce struct {
	Answer    string   `json:"answer"`
	Delay     int      `json:"delay"`
	Keyboard  Keyboard `json:"keyboard"`
	IsKb      bool     `json:"isKb"`
	IsNextMsg bool     `json:"isNextMsg"`
	Id        int      `json:"id"`
}
