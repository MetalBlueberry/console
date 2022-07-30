package console

import (
	"log"
	"os"
)

type Console struct {

	// counter
	timer Timer

	// counter
	counter Counter

	// logger
	LLog   *log.Logger
	LDebug *log.Logger
	LInfo  *log.Logger
	LWarn  *log.Logger
	LError *log.Logger
}

var DefaultConsole = &Console{

	// logger
	LLog:   log.New(os.Stdout, "", 0),
	LDebug: log.New(os.Stdout, "", 0),
	LInfo:  log.New(os.Stdout, "ℹ️  ", 0),
	LWarn:  log.New(os.Stdout, "⚠️  ", 0),
	LError: log.New(os.Stderr, "❌ ", 0),
}
