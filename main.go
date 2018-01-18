package main

import (
	"errors"
	"fmt"
	"os"
)

func showVersion() {
	fmt.Println(`blook 1.0.0 (Eclipse Public License 1.0, https://github.com/abtv/blook)`)
}

func showHelp() {
	fmt.Println(`Usage: blook from_pattern to_pattern file [file ...]
Makes binary search in one or more files. Returns all the lines which >= 'from_pattern' and <= 'to_pattern'.

Other options:
blook help - shows help
blook version - shows version`)
}

func filterFile(patternFrom string, patternTo string, filename string) {
	file, err := openFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer file.ptr.Close()

	start, err := blook(patternFrom, file.ptr, file.size, true)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	stop, err := blook(patternTo, file.ptr, file.size, false)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if start != -1 && start <= stop {
		maxBufferSize := int64(1024 * 1024)
		_, err = writeBytes(file.ptr, start, stop, os.Stdout, maxBufferSize)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Println("") //add newline symbol to the end of filter output
	}
}

func filter(patternFrom string, patternTo string, filenames []string) {
	for _, filename := range filenames {
		filterFile(patternFrom, patternTo, filename)
	}
}

func showError(err error) {
	fmt.Fprintln(os.Stderr, err)
}

func main() {
	cmdParams, err := getCmdParams()
	if err != nil {
		showError(err)
		return
	}

	switch cmdParams.command {
	case "version":
		showVersion()
	case "help":
		showHelp()
	case "filter":
		filter(cmdParams.patternFrom, cmdParams.patternTo, cmdParams.filenames)
	default:
		showError(errors.New("Unknown command. Please use 'blook help'"))
	}
}
