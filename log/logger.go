package log

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/Viva-Victoria/Lime-go/writer"
)

type Logger struct {
	formatter   Formatter
	writer      writer.Writer
	nonBlocking bool
}

func NewLogger(nonBlocking bool, w writer.Writer, f Formatter) *Logger {
	return &Logger{
		nonBlocking: nonBlocking,
		writer:      w,
		formatter:   f,
	}
}

func (l *Logger) write(event Event) error {
	data, err := l.formatter.Format(event)
	if err != nil {
		if err == ErrLevelDisabled {
			return nil
		}

		return err
	}

	if l.nonBlocking {
		l.writer.PostWrite(data)
		return nil
	}

	return l.writer.Write(data)
}

func (l *Logger) Entry(level Level, message string, err error) *Entry {
	pc, file, line, ok := runtime.Caller(3)
	funcName := ""
	if ok {
		funcInfo := runtime.FuncForPC(pc)
		funcName = funcInfo.Name()
		i := strings.LastIndex(funcName, "/")
		funcName = funcName[i+1:]
	}

	return &Entry{
		callback: l.write,
		Event: Event{
			Level:     level,
			Timestamp: time.Now(),
			Error:     err,
			Message:   message,
			Params:    make(map[string]interface{}),
			Caller: Caller{
				File:     file,
				Line:     line,
				Function: funcName,
			},
		},
	}
}

func (l *Logger) WriteLog(level Level, message string, err error, params ...interface{}) *Entry {
	return l.Entry(level, message, err).With(params...)
}

func (l *Logger) Writef(level Level, format string, err error, params ...interface{}) *Entry {
	return l.Entry(level, fmt.Sprintf(format, params...), err)
}

func (l *Logger) Trace(message string, params ...interface{}) {
	_ = l.WriteLog(LevelTrace, message, nil, params...).Submit()
}

func (l *Logger) Tracef(format string, params ...interface{}) *Entry {
	return l.Writef(LevelTrace, format, nil, params...)
}

func (l *Logger) Tracew(message string, params ...interface{}) error {
	return l.WriteLog(LevelTrace, message, nil, params...).Submit()
}

func (l *Logger) Debug(message string, params ...interface{}) {
	_ = l.WriteLog(LevelDebug, message, nil, params...).Submit()
}

func (l *Logger) Debugf(format string, params ...interface{}) *Entry {
	return l.Writef(LevelDebug, format, nil, params...)
}

func (l *Logger) Debugw(message string, params ...interface{}) error {
	return l.WriteLog(LevelDebug, message, nil, params...).Submit()
}

func (l *Logger) Info(message string, params ...interface{}) {
	_ = l.WriteLog(LevelInfo, message, nil, params...).Submit()
}

func (l *Logger) Infof(format string, params ...interface{}) *Entry {
	return l.Writef(LevelInfo, format, nil, params...)
}

func (l *Logger) Infow(message string, params ...interface{}) error {
	return l.WriteLog(LevelInfo, message, nil, params...).Submit()
}

func (l *Logger) Warn(message string, params ...interface{}) {
	_ = l.WriteLog(LevelWarn, message, nil, params...).Submit()
}

func (l *Logger) Warnf(format string, params ...interface{}) *Entry {
	return l.Writef(LevelWarn, format, nil, params...)
}

func (l *Logger) Warnw(message string, params ...interface{}) error {
	return l.WriteLog(LevelWarn, message, nil, params...).Submit()
}

func (l *Logger) Warne(message string, err error, params ...interface{}) {
	_ = l.WriteLog(LevelWarn, message, err, params...).Submit()
}

func (l *Logger) Warnef(format string, err error, params ...interface{}) *Entry {
	return l.Writef(LevelWarn, format, err, params...)
}

func (l *Logger) Warnew(message string, err error, params ...interface{}) error {
	return l.WriteLog(LevelWarn, message, err, params...).Submit()
}

func (l *Logger) Error(message string, err error, params ...interface{}) {
	_ = l.WriteLog(LevelError, message, err, params...).Submit()
}

func (l *Logger) Errorf(format string, err error, params ...interface{}) *Entry {
	return l.Writef(LevelError, format, err, params...)
}

func (l *Logger) Errorw(message string, err error, params ...interface{}) error {
	return l.WriteLog(LevelError, message, err, params...).Submit()
}

func (l *Logger) Panic(err error, message string, params ...interface{}) {
	_ = l.WriteLog(LevelPanic, message, err, params...).Submit()
	panic(fmt.Errorf("%s: %w", message, err))
}

func (l *Logger) Panicf(err error, format string, params ...interface{}) *Entry {
	entry := l.Writef(LevelPanic, format, err, params...)
	origin := entry.callback
	entry.callback = func(event Event) error {
		_ = origin(event)
		panic(fmt.Errorf("%s: %w", entry.Message, err))
	}
	return entry
}

func (l *Logger) Panicw(message string, err error, params ...interface{}) {
	_ = l.WriteLog(LevelPanic, message, err, params...).Submit()
	panic(err)
}

func (l *Logger) Write(bytes []byte) (int, error) {
	err := l.WriteLog(LevelTrace, string(bytes), nil).Submit()
	return len(bytes), err
}

func (l *Logger) Printf(format string, params ...interface{}) {
	_ = l.WriteLog(LevelTrace, fmt.Sprintf(format, params...), nil).Submit()
}
