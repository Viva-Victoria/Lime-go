package systemd

import (
	"fmt"
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

func writeIfNeeded(builder *StringBuilder, key, value string) {
	if len(value) > 0 && len(key) > 0 {
		builder.AppendStrings(" ", key, "=", value)
	}
}

func (f *Formatter) Format(event log.Event) ([]byte, error) {
	enabled, _, level := f.config.FormatLevel(event.Level)
	if !enabled {
		return nil, log.ErrLevelDisabled
	}

	builder := &StringBuilder{}

	_, timestamp := f.config.FormatTime(event.Timestamp)

	builder.AppendStrings(fmt.Sprintf("%s", timestamp), " ", level, " ", event.Message)

	if event.Error != nil {
		key, value := f.config.FormatError(event.Error)
		writeIfNeeded(builder, key, value)
	}

	key, value := f.config.FormatCaller(event.Caller)
	writeIfNeeded(builder, key, value)

	var v interface{}
	for key, v = range event.Params {
		writeIfNeeded(builder, key, f.config.FormatValue(v))
	}

	return builder.Bytes(), nil
}
