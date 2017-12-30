package main

import (
	"errors"
	"flag"
)

type CmdParams struct {
	pattern  string
	filename string
}

func getCmdParams() (CmdParams, error) {
	pattern := flag.String("pattern", "", "pattern prefix to search")
	filename := flag.String("file", "", "name of file")
	flag.Parse()

	if *pattern == "" {
		return CmdParams{}, errors.New("Error: 'pattern' is required parameter")
	}

	if *filename == "" {
		return CmdParams{}, errors.New("Error: 'file' is required parameter")
	}

	return CmdParams{pattern: *pattern, filename: *filename}, nil
}
