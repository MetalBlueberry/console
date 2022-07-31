package console

import (
	"log"
	"os"
	"sync"

	"github.com/olekukonko/tablewriter"
)

type Console struct {

	// table
	TableWriter func() *tablewriter.Table

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

	lMutex     sync.RWMutex
	groupDepth int
}

var DefaultLoggerFlags = log.LstdFlags | log.Lmsgprefix

var DefaultConsole = &Console{
	TableWriter: DefaultTableWriter,

	// logger
	LLog:   log.New(os.Stdout, "", DefaultLoggerFlags),
	LDebug: log.New(os.Stdout, "", DefaultLoggerFlags),
	LInfo:  log.New(os.Stdout, "ℹ️  ", DefaultLoggerFlags),
	LWarn:  log.New(os.Stdout, "⚠️  ", DefaultLoggerFlags),
	LError: log.New(os.Stderr, "❌ ", DefaultLoggerFlags),
}
