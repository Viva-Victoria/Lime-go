package systemd

import (
	"encoding/json"
	"github.com/Viva-Victoria/Lime-go/log"
)

type ValueFormatter func(value interface{}) string

type Config struct {
	log.Config
	FormatValue ValueFormatter
}

func DefaultConfig() Config {
	return Config{
		Config: log.Config{
			FormatTime:   log.DefaultTimeFormatter("timestamp"),
			FormatLevel:  log.DefaultLevelFormatter("level"),
			FormatCaller: log.DefaultCallerFormatter("stacktrace"),
			FormatError:  log.DefaultErrorFormatter("error"),
			MessageKey:   "message",
		},
		FormatValue: func(value interface{}) string {
			data, err := json.Marshal(value)
			if err != nil {
				return ""
			}

			return string(data)
		},
	}
}
