package main

import (
	"fmt"
	"os"
)

func showVersion() {
	fmt.Println(`blook 1.0.0 (Eclipse Public License 1.0, https://github.com/abtv/blook)`)
}

func showHelp() {
	fmt.Println(`Usage: blook from_pattern [-m to_pattern] [filename]
Makes binary search in [filename]. Returns all the lines which equals or more than 'from_pattern' to the end of the file.

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

	stop := file.size - 1
	if patternTo != "" {
		stop, err = blook(patternTo, file.ptr, file.size, false)
		stop++
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	if start != -1 && start <= stop {
		maxBufferSize := int64(1024 * 1024)
		_, err = writeBytes(file.ptr, start, stop, os.Stdout, maxBufferSize)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func filter(patternFrom string, patternTo string, filenames []string) {
	for _, filename := range filenames {
		filterFile(patternFrom, patternTo, filename)
	}
}

func showError() {
	fmt.Fprintln(os.Stderr, "Unknown parameter. Please use `blook help`")
}

func main() {
	cmdParams := getCmdParams()

	switch cmdParams.command {
	case "version":
		showVersion()
	case "help":
		showHelp()
	case "filter":
		filter(cmdParams.patternFrom, cmdParams.patternTo, cmdParams.filenames)
	case "error":
		showError()
	default:
		showError()
	}
}
