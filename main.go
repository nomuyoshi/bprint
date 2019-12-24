package main
import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	flag "github.com/spf13/pflag"
)
// options
var (
	tty      string
	theme    string
	listFlag bool
	helpFlag bool
)
var (
	exitCode int       = 0
	writer   io.Writer = os.Stdout
)
func init() {
	flag.BoolVarP(&helpFlag, "help", "h", false, "show this help message")
	flag.StringVar(&tty, "tty", "terminal256", `terminal type. "terminal256" or "terminal"`)
	flag.BoolVarP(&listFlag, "list", "l", false, "show themes list")
	flag.StringVarP(&theme, "theme", "t", "monokai", "theme")
}
func main() {
	bprintMain()
	os.Exit(exitCode)
}
func bprintMain() {
	flag.Parse()
	if helpFlag {
		usage()
		return
	}
	if listFlag {
		themeList()
		return
	}
	for _, path := range flag.Args() {
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			report(err)
			return
		}
		lexer := lexers.Match(path)
		if lexer == nil {
			lexer = lexers.Fallback
		}
		print(lexer, &contents)
	}
}
func print(lexer chroma.Lexer, contents *[]byte) {
	formatter := formatters.Get(tty)
	style := styles.Get(theme)
	iterator, err := lexer.Tokenise(nil, string(*contents))
	if err != nil {
		report(err)
		return
	}
	if err = formatter.Format(writer, style, iterator); err != nil {
		report(err)
		return
	}
}
func themeList() {
	fmt.Fprintln(writer, "Available themes list.")
	fmt.Fprintln(writer, "-----------------------")
	fmt.Fprintln(writer, strings.Join(styles.Names(), "\n"))
}
func usage() {
	fmt.Fprintln(writer, "usage: bprint [options] [filepath]")
	flag.PrintDefaults()
}
func report(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	exitCode = 2
}
