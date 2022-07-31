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

	time.Sleep(time.Millisecond * 500)
	console.Clear()
	console.Trace()

	console.Log("root")
	console.Group()
	console.Info("Wops!")
	console.GroupEnd()
	console.Log("perfect")
	console.Group("Found multiple errors")
	console.Error("error 1")
	console.Error("value %d was not expected", 10)
	console.Group("Valid values")
	console.Info(1,2,3,4,5)

}
