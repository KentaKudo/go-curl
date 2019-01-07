package main

import (
	"flag"
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

	v := fs.Bool("v", false, "Makes  curl  verbose  during the operation.")
	var X string
	fs.StringVar(&X, "X", "GET", "(HTTP) Specifies a custom request method to use when communicating with the HTTP server.")

	if err := fs.Parse(args); err != nil {
		return ExitCodeError
	}

	curl := &Curl{
		Method:  X,
		Verbose: *v,
	}
	curl.request(map[string]string{})
	return ExitCodeOK
}
