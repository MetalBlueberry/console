package main

import (
	"time"

	"github.com/metalblueberry/console"
)

func main() {
	console.Log("hello world")
	console.Info("this is an info message")
	console.Warn("this is a warn message")
	console.Error("this is an error message")

	console.Count()
	console.Count()
	console.Count("test")
	console.Count()
	console.Count("test")
	console.CountReset()
	console.Count()
	console.Count("test")
	console.CountReset("test")
	console.Count("test")

	console.Time()
	console.Time()
	time.Sleep(time.Millisecond * 100)
	console.TimeLog()
	console.Time("timer 2")
	console.Time("timer 2")
	time.Sleep(time.Millisecond * 100)
	console.TimeEnd()
	console.TimeLog("timer 2")
	console.TimeEnd("timer 2")
	console.TimeEnd()

	console.Trace()

}
