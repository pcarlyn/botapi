package models

type UpdateVariablesOK struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type Variable struct {
	ProfileId int    `json:"profile_id"`
	Name      string `json:"name"`
	Value     string `json:"value"`
}

type VariablesResponse struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
