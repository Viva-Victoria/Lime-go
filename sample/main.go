package main

import (
	"errors"
	"github.com/Viva-Victoria/Lime-go/log"
	"github.com/Viva-Victoria/Lime-go/log/json"
	"github.com/Viva-Victoria/Lime-go/stdout"
)

func main() {
	test()
}

func test() {
	w := stdout.NewSyncedCore()
	defer func() {
		_ = w.Sync()
	}()

	cfg := json.DefaultConfig()
	l := log.NewLogger(false, w, json.NewFormatter(cfg))
	l.Trace("hello world from Trace")

	err := l.Tracef("hello world from %s", "TraceF").With("name", "victoria").Submit()
	if err != nil {
		panic(err)
	}

	err = l.Tracew("hello world", "from", "TraceW")
	if err != nil {
		panic(err)
	}

	l.Warn("simple warning", "key", "value")
	err = l.Warnew("warning with error", errors.New("warning! error"))
	if err != nil {
		panic(err)
	}

	l.Panic(errors.New("FATAL! it is panic"), "panic error log")
}
