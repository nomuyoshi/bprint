package main

import (
	"fmt"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	flag "github.com/spf13/pflag"
	"io"
	"io/ioutil"
	"strings"
)

const (
	tty        string = "terminal256"
	exitCodeOK int    = iota
	exitCodeErr
)

// options
var (
	theme string
	list  bool
	help  bool
)

// CLI は出力先を管理するStruct
type CLI struct {
	outStream, errStream io.Writer
}

// Run bprintのメイン処理
func (c *CLI) Run(args []string) int {
	flags := flag.NewFlagSet("bprint", flag.ContinueOnError)
	flags.BoolVarP(&help, "help", "h", false, "show this help message")
	flags.BoolVarP(&list, "list", "l", false, "show themes list")
	flags.StringVarP(&theme, "theme", "t", "monokai", "theme")
	flags.SetOutput(c.outStream)
	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprintf(c.errStream, "illegal option exits.\n")
		flags.PrintDefaults()
		return exitCodeErr
	}
	if help {
		usage(c, flags)
		return exitCodeOK
	}
	if list {
		themes(c)
		return exitCodeOK
	}
	for _, path := range flags.Args() {
		if err := print(c, path); err != nil {
			fmt.Fprintf(c.errStream, "failed to print %s.\n", path)
			return exitCodeErr
		}
	}
	return exitCodeOK
}
func print(c *CLI, path string) error {
	formatter := formatters.Get(tty)
	style := styles.Get(theme)
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	lexer := lexers.Match(path)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	iterator, err := lexer.Tokenise(nil, string(contents))
	if err != nil {
		return err
	}
	if err = formatter.Format(c.outStream, style, iterator); err != nil {
		return err
	}
	return nil
}
func themes(c *CLI) {
	fmt.Fprintln(c.outStream, "Available themes list.\n-----------------------")
	fmt.Fprintln(c.outStream, strings.Join(styles.Names(), "\n"))
}
func usage(c *CLI, flags *flag.FlagSet) {
	fmt.Fprintln(c.outStream, "usage: bprint [options] [filepath]")
	flags.PrintDefaults()
}
