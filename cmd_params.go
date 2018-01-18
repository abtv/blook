package main

import (
	"errors"
	"os"
)

type CmdParams struct {
	command     string
	patternFrom string
	patternTo   string
	filenames   []string
}

func getCmdParams() (CmdParams, error) {
	args := os.Args[1:]
	if len(args) == 0 {
		return CmdParams{command: "help"}, nil
	} else if len(args) == 1 {
		return CmdParams{command: args[0]}, nil
	} else if len(args) == 2 {
		return CmdParams{}, errors.New("you need to provide at least one file")
	} else {
		return CmdParams{command: "filter", patternFrom: args[0], patternTo: args[1], filenames: args[2:]}, nil
	}
}
