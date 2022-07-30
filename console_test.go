package console_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/metalblueberry/console"
)

type Simple struct {
	Name string
}

func TestConsole_Log(t *testing.T) {
	type args struct {
		any []any
	}
	tests := []struct {
		args   []any
		expect string
	}{
		{
			args:   []any{"hello world"},
			expect: "hello world\n",
		},
		{
			args: []any{Simple{
				Name: "hello world",
			}},
			expect: "console_test.Simple{Name:\"hello world\"}\n",
		},
		{
			args:   []any{"number", 1, "is number one"},
			expect: "number 1 is number one\n",
		},
		{
			args:   []any{"object %o is simple", Simple{Name: "hello world"}},
			expect: "object console_test.Simple{Name:\"hello world\"} is simple\n",
		},
		{
			args:   []any{"house %s contains %d humans", "Adams", 0},
			expect: "house Adams contains 0 humans\n",
		},
		{
			args:   []any{"Oxygen level is %.3f", 5.123},
			expect: "Oxygen level is 5.123\n",
		},
		{
			args:   []any{"this has", 0, "substitutions strings"},
			expect: "this has 0 substitutions strings\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.expect, func(t *testing.T) {
			buf := &bytes.Buffer{}
			c := &console.Console{
				LLog:   log.New(buf, "", 0),
				LInfo:  log.New(buf, "Info", 0),
				LWarn:  log.New(buf, "Warn", 0),
				LError: log.New(buf, "Error", 0),
			}
			c.Log(tt.args...)

			actual := buf.String()
			if actual != tt.expect {
				t.Logf("expected \"%s\", actual \"%s\"", tt.expect, actual)
				t.Fail()
			}
		})
	}
}
