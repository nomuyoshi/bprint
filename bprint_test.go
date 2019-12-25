package main

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("--list option", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		cli := &CLI{outStream: buffer, errStream: buffer}

		args := strings.Split("bprint --list", " ")
		golden := "./testdata/themes.golden"
		want, err := ioutil.ReadFile(golden)
		if err != nil {
			t.Fatalf("ioutil.ReadFile failed. %s", err)
		}
		r := cli.Run(args)
		got := buffer.Bytes()
		if !bytes.Equal(got, want) {
			t.Errorf("unexpected output.\ngot: %s", buffer.String())
		}

		if r != 0 {
			t.Errorf("unexpected return value.\ngot: %q", r)
		}
	})

	t.Run("--help option", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		cli := &CLI{outStream: buffer, errStream: buffer}

		args := strings.Split("bprint --help", " ")
		want := `usage: bprint [options] [filepath]
  -h, --help           show this help message
  -l, --list           show themes list
  -t, --theme string   theme (default "monokai")
      --tty string     terminal type. "terminal256" or "terminal" (default "terminal256")
`
		r := cli.Run(args)
		got := buffer.String()

		if got != want {
			t.Errorf("unexpected output.\ngot: %s", got)
		}
		if r != 0 {
			t.Errorf("unexpected return value.\ngot: %q", r)
		}
	})
}
