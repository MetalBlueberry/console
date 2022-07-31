package console

import (
	"os"
	"os/exec"
	"runtime"
)

func Clear() {
	DefaultConsole.Clear()
}

func (c *Console) Clear() {
	cmd := getClearCommand()
	if cmd == nil {
		c.Error("Platform %s is unsupported! I can't clear terminal screen :(", runtime.GOOS)
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getClearCommand() *exec.Cmd {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("clear")
	case "windows":
		return exec.Command("cmd", "/c", "cls")
	default:
		return nil
	}
}
