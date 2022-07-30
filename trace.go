package console

import (
	"runtime"
)

func Trace() { DefaultConsole.Trace() }

func (c *Console) Trace() {
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	c.Log(string(buf))
}
