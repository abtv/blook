package main

import (
	"log"
	"os"
)

var maxBufferSize = int64(1024 * 1024)

func filterFile(pattern string, filename string) {
	file, err := openFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.ptr.Close()

	start, err := filterLines(pattern, file)
	if err != nil {
		log.Fatal(err)
	}

	if start != -1 && start < file.size {
		_, err = writeBytes(file.ptr, start, file.size, os.Stdout, maxBufferSize)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	cmdParams, err := getCmdParams()
	if err != nil {
		log.Fatal(err)
	}

	for _, filename := range cmdParams.filenames {
		filterFile(cmdParams.pattern, filename)
	}
}
