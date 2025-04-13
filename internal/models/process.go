package models

import "time"

type Process struct {
	Path string   `json:"path"`
	Args []string `json:"args"`
}

type ProcessAnswer struct {
	Command Process   `json:"cmd"`
	Pid     uint32    `json:"pid"`
	RunAt   time.Time `json:"run_at"`
}
