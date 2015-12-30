package main

import (
	"os"
)

var VERSION string = `0.1.0`

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
