package json

import (
	"encoding/json"

	"github.com/Viva-Victoria/Lime-go/log"
)

type Marshaller func(indent string, value interface{}) ([]byte, error)

type Config struct {
	log.Config
	Marshal Marshaller
	Indent  string
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
		Indent: "  ",
		Marshal: func(indent string, value interface{}) ([]byte, error) {
			if len(indent) > 0 {
				return json.MarshalIndent(value, "", indent)
			}

			return json.Marshal(value)
		},
	}
}
