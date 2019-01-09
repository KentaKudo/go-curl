package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
	ExitCodeFileError
)

type CLI struct {
	outStream, errStream io.Writer
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args[1:]))
}

func (c *CLI) Run(args []string) int {
	fs := flag.NewFlagSet("curl", flag.ContinueOnError)
	fs.SetOutput(c.errStream)

	var v bool
	fs.BoolVar(&v, "v", false, "Makes  curl  verbose  during the operation.")
	var X string
	fs.StringVar(&X, "X", "GET", "(HTTP) Specifies a custom request method to use when communicating with the HTTP server.")
	var headers Headers
	fs.Var(&headers, "header", "(HTTP) Extra header to include in the request when sending HTTP to a server.")
	var H Headers
	fs.Var(&H, "H", "(HTTP) Extra header to include in the request when sending HTTP to a server.")
	var d string
	fs.StringVar(&d, "d", "", "(HTTP) Sends the specified data in a POST request to the HTTP server.")

	if err := fs.Parse(args); err != nil {
		return ExitCodeError
	}

	url := fs.Arg(0)
	if url == "" {
		fmt.Fprintf(c.errStream, "no URL specified!\n")
		fs.PrintDefaults()
		return ExitCodeError
	}

	curl := NewCurl(url, v, append(headers, H...))

	var r string
	var err error
	switch X {
	case "GET":
		r, err = curl.Get(map[string]string{})
	case "POST":
		r, err = curl.Post(d)
	default:
		fmt.Fprintf(c.errStream, "Unknown Method: %s", X)
		return ExitCodeError
	}

	if err != nil {
		fmt.Fprintf(c.errStream, "Error on requesting %s: %s", url, err)
	}

	fmt.Fprint(c.outStream, r)

	return ExitCodeOK
}
