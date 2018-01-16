package main

import (
	"os"
)

type CmdParams struct {
	command     string
	patternFrom string
	patternTo   string
	filenames   []string
}

func getCmdParams() CmdParams {
	args := os.Args[1:]
	if len(args) == 0 {
		return CmdParams{command: "help"}
	} else if len(args) == 1 {
		return CmdParams{command: args[0]}
	} else if len(args) == 2 {
		if args[1] == "-m" {
			return CmdParams{command: "error"}
		}
		return CmdParams{command: "filter", patternFrom: args[0], patternTo: "", filenames: args[1:]}
	} else {
		if args[1] == "-m" {
			if len(args) == 3 {
				return CmdParams{command: "error"}
			}
			return CmdParams{command: "filter", patternFrom: args[0], patternTo: args[2], filenames: args[3:]}
		}
		return CmdParams{command: "filter", patternFrom: args[0], patternTo: "", filenames: args[1:]}
	}
}
