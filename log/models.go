package log

import "time"

type Caller struct {
	File     string
	Function string
	Line     int
}

type Event struct {
	Params    map[string]interface{}
	Message   string
	Error     error
	Timestamp time.Time
	Caller    Caller
	Level     Level
}

type Entry struct {
	callback func(event Event) error
	Event
}

func (e *Entry) With(params ...interface{}) *Entry {
	if len(params)%2 != 0 {
		return e
	}

	for i := 0; i < len(params); i += 2 {
		key, ok := params[i].(string)
		if !ok {
			continue
		}

		e.Params[key] = params[i+1]
	}

	return e
}

func (e *Entry) Submit() error {
	return e.callback(e.Event)
}
