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
		cli.Run(args)
		got := buffer.Bytes()
		if !bytes.Equal(got, want) {
			t.Errorf("themes() output does not match .golden file")
		}
	})
}
