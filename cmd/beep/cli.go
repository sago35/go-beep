package main

import (
	"flag"
	"fmt"
	"github.com/sago35/go-beep"
	"io"
)

const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
	ExitCodeError
)

type CLI struct {
	outStream io.Writer
	errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	var version bool
	var score string

	flags := flag.NewFlagSet("beep", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.BoolVar(&version, "version", false, "Print version information and quit")
	flags.StringVar(&score, "score", ``, "Score")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	if version {
		fmt.Fprintf(c.errStream, "beep version %s\n", VERSION)
		return ExitCodeOK
	}

	if len(score) > 0 {
		beep.Play(score)
	}

	return ExitCodeOK
}
