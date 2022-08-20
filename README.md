# Console

Console package allows you to use `console.log()` and other `console` commands from the javascript ecosystem in go.

It is designed to feel like javascript but to still be compatible with go tooling

It is a wrapper over the `log` package with some fancy stuff from javascript such as log groups.

## Examples

```go
package main

import (
	"github.com/metalblueberry/console"
)

func main() {
	console.Log("hello world")
	console.Info("this is an info message")
	console.Warn("this is a warn message")
	console.Error("this is an error message")
}
```

## Implementation list

I accept contributions

- [x] assert
- [x] clear
- [x] count
- [x] countReset
- [x] debug
- [ ] ~~dir~~ Cannot print interactive info
- [ ] ~~dirxml~~ Cannot print interactive info
- [x] error
- [x] group
- [ ] ~~groupCollapsed~~ cannot collapse on a terminal
- [x] groupEnd
- [x] info
- [x] log
- [ ] table - It is just complicated :harold:
- [x] time
- [x] timeEnd
- [x] timeLog
- [x] trace
- [x] warn

## Why?

I've found myself writing `console.log` so many times that I've decided that I don't want to fight against my instincts any more. I hate javascript but I can't get rid of `console.log`.
