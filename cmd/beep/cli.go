package main

import (
	"fmt"
	"github.com/sago35/go-beep"
	"gopkg.in/alecthomas/kingpin.v2"
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

var (
	version = kingpin.Flag("version", "Print version information and quit").Bool()
	score   = kingpin.Arg("score", "Input score").String()
)

func (c *CLI) Run(args []string) int {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	if *version {
		fmt.Fprintf(c.errStream, "beep version %s\n", VERSION)
		return ExitCodeOK
	}

	if len(*score) > 0 {
		beep.Play(*score, 120)
	}

	return ExitCodeOK
}
