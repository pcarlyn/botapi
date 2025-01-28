package models

type PostRequestProfile struct {
	Id       int      `json:"id"`
	Isadmin  bool     `json:"is"`
	Statuses []string `json:"statuses"`
	Achives  []string `json:"achives"`
}

type ResponseProfile struct {
	Id         int      `json:"id"`
	Active     bool     `json:"active"`
	Registered string   `json:"registered"`
	Statuses   []string `json:"statuses"`
	Last_Visit string   `json:"last_visit"`
	Isadmin    bool     `json:"is"`
	Achives    []string `json:"achives"`
}

type GetResponseProfile struct {
	Error   bool              `json:"error"`
	Message string            `json:"message"`
	Data    []ResponseProfile `json:"data"`
}

type LVUpdate struct {
	Id int `json:"id"`
}

type UpdateResponse struct {
	Message string `json:"message"`
}
