package log

import (
	"time"
)

type Caller struct {
	File     string
	Function string
	Line     int
}

type Event struct {
	Level     Level
	Timestamp time.Time
	Caller    Caller
	Error     error
	Message   string
	Params    map[string]interface{}
}
