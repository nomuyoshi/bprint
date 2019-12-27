package main

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("Failed: illegal option error.", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		cli := &CLI{outStream: buffer, errStream: buffer}
		args := strings.Split("bprint --illegal main.go", " ")
		want := `illegal option exits.
  -h, --help           show this help message
  -l, --list           show themes list
  -t, --theme string   theme (default "monokai")
`
		code := cli.Run(args)
		got := buffer.String()
		if want != got {
			t.Errorf("unexpected output.\nwant: %s\ngot: %s", want, got)
		}
		if code != exitCodeErr {
			t.Errorf("unexpected exit code. want: %d, got: %d", exitCodeErr, code)
		}
	})
	t.Run("Success: list option", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		cli := &CLI{outStream: buffer, errStream: buffer}
		args := strings.Split("bprint --list", " ")
		golden := "./testdata/themes.golden"
		want, err := ioutil.ReadFile(golden)
		if err != nil {
			t.Fatalf("ioutil.ReadFile failed. %s", err)
		}
		code := cli.Run(args)
		got := buffer.Bytes()
		if !bytes.Equal(got, want) {
			t.Errorf("unexpected output.\nwant: %s\ngot: %s", want, got)
		}
		if code != exitCodeOK {
			t.Errorf("unexpected exit code. want: %d, got: %d", exitCodeOK, code)
		}
	})
	t.Run("Success: help option", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		cli := &CLI{outStream: buffer, errStream: buffer}
		args := strings.Split("bprint --help", " ")
		want := `usage: bprint [options] [filepath]
  -h, --help           show this help message
  -l, --list           show themes list
  -t, --theme string   theme (default "monokai")
`
		code := cli.Run(args)
		got := buffer.String()
		if want != got {
			t.Errorf("unexpected output.\nwant: %s\ngot: %s", want, got)
		}
		if code != exitCodeOK {
			t.Errorf("unexpected exit code. want: %d, got: %d", exitCodeOK, code)
		}
	})
	t.Run("Success: print file contents.", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		cli := &CLI{outStream: buffer, errStream: buffer}
		args := strings.Split("bprint ./testdata/template.go", " ")
		want := "[38;5;197mpackage[0m[38;5;231m [0m[38;5;148mmain[0m[38;5;231m"
		code := cli.Run(args)
		got := buffer.String()
		if !strings.HasPrefix(got, want) {
			t.Errorf("unexpected output.\nwant: %s\ngot: %s", want, got)
		}
		if code != exitCodeOK {
			t.Errorf("unexpected exit code. want: %d, got: %d", exitCodeOK, code)
		}
	})
	t.Run("Success: print file contents with theme option.", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		cli := &CLI{outStream: buffer, errStream: buffer}
		args := strings.Split("bprint -t=github ./testdata/template.go", " ")
		want := "[1m[38;5;16mpackage[0m main"
		code := cli.Run(args)
		got := buffer.String()
		if !strings.HasPrefix(got, want) {
			t.Errorf("unexpected output.\nwant: %s\ngot: %s", want, got)
		}
		if code != exitCodeOK {
			t.Errorf("unexpected exit code. want: %d, got: %d", exitCodeOK, code)
		}
	})
	t.Run("Success: print unsupported file contents.", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		cli := &CLI{outStream: buffer, errStream: buffer}
		args := strings.Split("bprint ./testdata/template.test", " ")
		want := "[38;5;231mpackage main[0m[38;5;231m"
		code := cli.Run(args)
		got := buffer.String()
		if !strings.HasPrefix(got, want) {
			t.Errorf("unexpected output.\nwant: %s\ngot: %s", want, got)
		}
		if code != exitCodeOK {
			t.Errorf("unexpected exit code. want: %d, got: %d", exitCodeOK, code)
		}
	})
	t.Run("Failed: file not found.", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		cli := &CLI{outStream: buffer, errStream: buffer}
		path := "./testdata/not_exist_file.go"
		args := strings.Split("bprint "+path, " ")
		code := cli.Run(args)
		want := "failed to print " + path + ".\n"
		got := buffer.String()
		if want != got {
			t.Errorf("unexpected output.\nwant: %s\ngot: %s", want, got)
		}
		if code != exitCodeErr {
			t.Errorf("unexpected exit code. want: %d, got: %d", exitCodeErr, code)
		}
	})
}
