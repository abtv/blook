package main

import (
	"os"
)

type File struct {
	ptr  *os.File
	size int64
}

func openFile(filename string) (File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return File{}, err
	}

	fi, err := file.Stat()
	if err != nil {
		file.Close()
		return File{}, err
	}
	size := fi.Size()

	return File{ptr: file, size: size}, nil
}
