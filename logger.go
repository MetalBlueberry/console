package console

import (
	"fmt"
	"regexp"
	"strings"
)

func Log(any ...any)   { DefaultConsole.Log(any...) }
func Info(any ...any)  { DefaultConsole.Info(any...) }
func Warn(any ...any)  { DefaultConsole.Warn(any...) }
func Error(any ...any) { DefaultConsole.Error(any...) }

func (c *Console) Log(any ...any) {
	c.LLog.Print(c.prepare(any...))
}

func (c *Console) Debug(any ...any) {
	c.LDebug.Print(c.prepare(any...))
}

func (c *Console) Info(any ...any) {
	c.LInfo.Print(c.prepare(any...))
}

func (c *Console) Warn(any ...any) {
	c.LWarn.Print(c.prepare(any...))
}

func (c *Console) Error(any ...any) {
	c.LError.Print(c.prepare(any...))
}

func (c *Console) prepare(any ...any) string {
	if len(any) == 0 {
		return ""
	}

	if c.hasSubstitutions(any[0]) {
		return c.logSubstitutions(any[0].(string), any[1:]...)
	}

	line := &strings.Builder{}
	for i := range any {
		l := c.logAny(any[i])
		line.WriteString(l)
		if i < len(any)-1 {
			line.WriteString(" ")
		}
	}
	return line.String()
}

func (c *Console) logAny(obj any) string {
	switch v := obj.(type) {
	case string:
		return c.logStr(v)
	default:
		return c.logObj(v)
	}
}

func (c *Console) logObj(obj any) string {
	return fmt.Sprintf("%#v", obj)
}

func (c *Console) logStr(msg string) string {
	return fmt.Sprintf(msg)
}

var floatSubstitutionRegex = regexp.MustCompile(`^[^%]*%\d*?[\.]{0,1}\d*?f`)

func (c *Console) hasSubstitutions(msg any) bool {
	v, ok := msg.(string)
	if !ok {
		return false
	}
	substitutions := []string{
		"%o", "%O", "%s", "%d", "%i",
	}

	for _, s := range substitutions {
		if strings.Contains(v, s) {
			return true
		}
	}
	m := floatSubstitutionRegex.MatchString(v)
	if m {
		return true
	}

	return false
}
func (c *Console) mapSubstitutions(template string) string {
	maps := map[string]string{
		"%o": "%#v",
		"%O": "%#v",
		"%i": "%d",
	}

	for a, b := range maps {
		template = strings.ReplaceAll(template, a, b)
	}
	return template
}

func (c *Console) logSubstitutions(template string, obj ...any) string {
	return fmt.Sprintf(c.mapSubstitutions(template), obj...)
}
