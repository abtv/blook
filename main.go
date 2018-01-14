package main

import (
	"fmt"
	"os"
)

func showVersion() {
	fmt.Println(`blook 1.0.0 (Eclipse Public License 1.0, https://github.com/abtv/blook)`)
}

func showHelp() {
	fmt.Println(`Usage: blook from_pattern [filename]
Makes binary search in [filename]. Returns all the lines which equals or more than 'from_pattern' to the end of the file.

Other options:
blook help - shows help
blook version - shows version`)
}

func filterFile(pattern string, filename string) {
	file, err := openFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer file.ptr.Close()

	start, err := filterLines(pattern, file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if start != -1 && start < file.size {
		maxBufferSize := int64(1024 * 1024)
		_, err = writeBytes(file.ptr, start, file.size, os.Stdout, maxBufferSize)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func filter(pattern string, filenames []string) {
	for _, filename := range filenames {
		filterFile(pattern, filename)
	}
}

func main() {
	cmdParams := getCmdParams()

	switch cmdParams.command {
	case "version":
		showVersion()
	case "help":
		showHelp()
	case "filter":
		filter(cmdParams.pattern, cmdParams.filenames)
	default:
		fmt.Fprintln(os.Stderr, "Unknown parameter. Please use `blook help`")
	}
}
