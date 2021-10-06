package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/theovassiliou/base64url"
	exitcodes "github.com/theovassiliou/go-exitcodes"

	"github.com/jpillora/opts"
)

/*
base64url -- Encode and decode using Base64Url representation

DESCRIPTION
     base64url encodes and decodes Base64URL data, as specified in RFC 4648.  With
     no options, base64url reads raw data from stdin and writes encoded data as a
     continuous block to stdout.

Usage:	base64url [-hvDd] [-b num] [-i in_file] [-o out_file]
  -h, --help     display this message
  -Dd, --decode   decodes input
  -b, --break    break encoded string into num character lines
  -i, --input    input file (default: "-" for stdin)
  -o, --output   output file (default: "-" for stdout)
*/

var conf = config{}

//set this via ldflags (see https://stackoverflow.com/q/11354518)
const pVersion = ".1"

// version is the current version number as tagged via git tag 1.0.0 -m 'A message'
var (
	version  = "0.1" + pVersion + "-src"
	commit   string
	branch   string
	repoName string = "github.com/theovassiliou/base64url"
)

type config struct {
	Decode bool   `help:"decodes input"`
	Input  string `help:"input file (\"-\" for stdin)"`
	Output string `help:"output file (\"-\" for stdout)"`
	Break  int    `help:"Insert line breaks every count characters.  Default is 0, which generates an unbroken stream"`
}

func main() {
	conf = config{
		Decode: false,
		Input:  "-",
		Output: "-",
		Break:  0,
	}

	//parse config
	opts.New(&conf).
		Repo(repoName).
		Version(FormatFullVersion("base64url", version, branch, commit)).
		Parse()

	inputFile := os.Stdin
	if conf.Input != "-" {
		i, err := os.Open(conf.Input)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(exitcodes.CANT_OPEN_INPUT)
		}
		inputFile = i
	}

	outputFile := os.Stdout

	if conf.Output != "-" {
		i, err := os.Create(conf.Output)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(exitcodes.CANT_CREATE_OUTPUT_FILE)
		}
		outputFile = i
	}

	in := bufio.NewReader(inputFile)
	buf := new(bytes.Buffer)
	buf.ReadFrom(in)

	if conf.Decode {
		n := normalizeNewlines(buf.String())
		output, err := base64url.Decode(strings.ReplaceAll(n, "\n", ""))
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(exitcodes.DATA_FORMAT_ERROR)
		}
		outputFile.Write(output)
	} else {
		output := base64url.Encode(buf.Bytes())
		if conf.Break != 0 {
			outputFile.Write([]byte(hard_wrap(output, conf.Break)))
		} else {
			outputFile.Write([]byte(output))
		}
	}
}

// taken from https://gist.github.com/kennwhite/f3881b815f43e0d9d7bd3ef8166e5d1b
func hard_wrap(text string, colBreak int) string {
	if colBreak < 1 {
		return text
	}

	text = strings.TrimSpace(text)
	wrapped := ""

	var i int
	for i = 0; len(text[i:]) > colBreak; i += colBreak {
		wrapped += text[i:i+colBreak] + "\n"
	}
	wrapped += text[i:]

	return wrapped
}

// Inspired from from https://www.programming-books.io/essential/go/normalize-newlines-1d3abcf6f17c4186bb9617fa14074e48
func normalizeNewlines(d string) string {
	// replace CR LF \r\n (windows) with LF \n (unix)
	d = strings.ReplaceAll(d, "\r\n", "\n")
	// replace CF \r (mac) with LF \n (unix)
	d = strings.ReplaceAll(d, "\r", "\n")
	return d
}

// FormatFullVersion formats for a cmdName the version number based on version, branch and commit
func FormatFullVersion(cmdName, version, branch, commit string) string {
	var parts = []string{cmdName}

	if version != "" {
		parts = append(parts, version)
	} else {
		parts = append(parts, "unknown")
	}

	if branch != "" || commit != "" {
		if branch == "" {
			branch = "unknown"
		}
		if commit == "" {
			commit = "unknown"
		}
		git := fmt.Sprintf("(git: %s %s)", branch, commit)
		parts = append(parts, git)
	}

	return strings.Join(parts, " ")
}
