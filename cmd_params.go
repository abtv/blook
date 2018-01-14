package main

import (
	"os"
)

type CmdParams struct {
	command   string
	pattern   string
	filenames []string
}

func getCmdParams() CmdParams {
	args := os.Args[1:]
	if len(args) == 0 {
		return CmdParams{command: "help"}
	} else if len(args) == 1 {
		return CmdParams{command: args[0]}
	} else {
		return CmdParams{command: "filter", pattern: args[0], filenames: args[1:]}
	}
}
