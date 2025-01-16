package models

type GetResponseAnswer struct {
	Error   bool           `json:"error"`
	Message string         `json:"message"`
	Data    ResponseAnswer `json:"data"`
}

type ResponseAnswer struct {
	Id            int         `json:"id"`
	Answer        string      `json:"answer"`
	IsKb          bool        `json:"isKb"`
	Keyboard      Keyboard    `json:"keyboard"`
	Conditions    []Condition `json:"conditions"`
	SetVariable   string      `json:"set_variable"`
	SetValue      string      `json:"set_value"`
	State         string      `json:"state"`
	NextState     string      `json:"nextState"`
	Delay         int         `json:"delay"`
	IsNextMessage bool        `json:"isNextMsg"`
	NextMessage   int         `json:"nextMsg"`
}

type Condition struct {
	Caption   string `json:"caption"`
	Variable  string `json:"variable"`
	Operation string `json:"operation"`
	Value     string `json:"value"`
}

type Keyboard struct {
	Type    string   `json:"type"`
	Buttons []Button `json:"buttons"`
}

type Button struct {
	Caption string `json:"caption"`
	Data    string `json:"data"`
	Row     int    `json:"row"`
	Order   int    `json:"order"`
}
