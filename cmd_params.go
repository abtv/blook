package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type CmdParams struct {
	command   string
	pattern   string
	filenames []string
}

func getCmdParams() (CmdParams, error) {
	args := os.Args[1:]

	fmt.Println(args)

	pattern := flag.String("pattern", "", "pattern prefix to search")
	filename := flag.String("file", "", "name of file")
	flag.Parse()

	if *pattern == "" {
		return CmdParams{}, errors.New("Error: 'pattern' is required parameter")
	}

	if *filename == "" {
		return CmdParams{}, errors.New("Error: 'file' is required parameter")
	}

	return CmdParams{pattern: *pattern, filenames: []string{*filename}}, nil
}
