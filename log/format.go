package log

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrLevelDisabled = errors.New("level disabled")
)

type Formatter interface {
	Format(event Event) ([]byte, error)
}

type LevelFormatter func(level Level) (enabled bool, key string, value string)

func DefaultLevelFormatter(key string, disable ...Level) LevelFormatter {
	disabledMap := make(map[Level]bool)
	for _, l := range disable {
		disabledMap[l] = true
	}

	return func(level Level) (bool, string, string) {
		_, ok := disabledMap[level]
		if ok {
			return false, "", ""
		}

		switch level {
		case LevelTrace:
			return true, key, "TRACE"
		case LevelDebug:
			return true, key, "DEBUG"
		case LevelInfo:
			return true, key, "INFO"
		case LevelWarn:
			return true, key, "WARN"
		case LevelError:
			return true, key, "ERROR"
		case LevelPanic:
			return true, key, "PANIC"
		default:
			return false, "", ""
		}
	}
}

type TimeFormatter func(time time.Time) (key string, value interface{})

func DefaultTimeFormatter(key string) TimeFormatter {
	return func(timestamp time.Time) (string, interface{}) {
		return key, timestamp.Format(time.RFC3339)
	}
}

type ErrorFormatter func(e error) (key string, value string)

func DefaultErrorFormatter(key string) ErrorFormatter {
	return func(e error) (string, string) {
		return key, e.Error()
	}
}

type CallerFormatter func(caller Caller) (key string, value string)

func DefaultCallerFormatter(key string) CallerFormatter {
	return func(caller Caller) (string, string) {
		return key, fmt.Sprintf("%s:%d %s", caller.File, caller.Line, caller.Function)
	}
}

type Config struct {
	FormatTime   TimeFormatter
	FormatLevel  LevelFormatter
	FormatError  ErrorFormatter
	FormatCaller CallerFormatter
	MessageKey   string
}
