package json

import (
	"github.com/Viva-Victoria/Lime-go/log"
)

type Formatter struct {
	config Config
}

func NewFormatter(config Config) *Formatter {
	return &Formatter{
		config: config,
	}
}

func (f *Formatter) Format(event log.Event) ([]byte, error) {
	enabled, key, value := f.config.FormatLevel(event.Level)
	if !enabled {
		return nil, log.ErrLevelDisabled
	}

	data := make(map[string]interface{})
	data[key] = value

	key, v := f.config.FormatTime(event.Timestamp)
	data[key] = v

	key, value = f.config.FormatCaller(event.Caller)
	data[key] = value

	if event.Error != nil {
		key, value = f.config.FormatError(event.Error)
		data[key] = value
	}

	data[f.config.MessageKey] = event.Message

	for key, v = range event.Params {
		data[key] = v
	}

	return f.config.Marshal(f.config.Indent, data)
}
