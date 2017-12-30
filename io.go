package main

import (
	"errors"
	"os"
)

func minimum(size int64, maxBufferSize int) int {
	if size > int64(maxBufferSize) {
		return maxBufferSize
	}
	return int(size)
}

func writeBytes(file *os.File, start int64, stop int64) (int64, error) {
	var bytesWritten int64
	bytesWritten = 0
	if start >= stop {
		return bytesWritten, nil
	}

	maxBufferSize := 1024 * 1024
	file.Seek(start, 0)
	buffer := make([]byte, minimum(stop-start, maxBufferSize))
	for current := start; current < stop; {
		bufferSize := minimum(stop-current, maxBufferSize)
		if bufferSize < maxBufferSize {
			buffer = make([]byte, bufferSize)
		}

		n, err := file.Read(buffer)
		if err != nil {
			return bytesWritten, err
		} else if n < bufferSize {
			return bytesWritten, errors.New("Error: unexpected end of input")
		}
		n, err = os.Stdout.Write(buffer)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += int64(n)

		current += int64(bufferSize)
	}

	return bytesWritten, nil
}

// newLineIndex returns index of newline symbol in buffer;
// if no newline symbol found returns -1
func newLineIndex(buffer []byte, diff int) int {
	n := len(buffer)
	if n == 0 {
		return -1
	}

	idx := 0
	if diff == -1 {
		idx = n - 1
	}

	for {
		if n == 0 {
			return -1
		}

		if buffer[idx] == '\n' {
			return idx
		}
		idx = idx + diff
		n--
	}
}

// findBorder searches for newline symbol in [from; to]
// when diff = 1 makes forward search (`from` -> `to`)
// when diff = -1 makes backward search (`to` -> `from`)
func findBorder(filePtr *os.File, from int64, to int64, diff int, maxBufferSize int) (int64, error) {
	size := to - from + int64(1)
	currentSize := minimum(size, maxBufferSize)

	position := from
	if diff == -1 {
		position = to - int64(currentSize)
	}
	buffer := make([]byte, currentSize)

	for {
		if size == 0 {
			return -1, nil
		}
		if len(buffer) != currentSize {
			buffer = make([]byte, currentSize)
		}

		filePtr.Seek(position, 0)
		n, err := filePtr.Read(buffer)
		if err != nil {
			return -1, err
		} else if n < currentSize {
			return -1, errors.New("Error: unexpected end of input")
		}

		idx := newLineIndex(buffer, diff)
		if idx >= 0 {
			return position + int64(idx), nil
		}

		position = position + int64(diff*currentSize)
		size = size - int64(currentSize)
		currentSize = minimum(size, maxBufferSize)
	}
}

// findString searches string borders
// returns (leftBorder, rightBorder, error)
func findString(file *os.File, from int64, to int64) (int64, int64, error) {
	//TODO naive implementation; need to write unit tests first & fix implementation mistakes
	maxBufferSize := 64 * 1024
	middle := (from + to) / 2
	strFrom, err := findBorder(file, middle, from, -1, maxBufferSize)
	if err != nil {
		return -1, -1, err
	}
	strTo, err := findBorder(file, middle+1, to, 1, maxBufferSize)
	if err != nil {
		return -1, -1, err
	}
	return strFrom, strTo, nil
}

// getString returns string from `file` in [from; to]
func getString(file *os.File, from int64, to int64) (string, error) {
	bufferSize := to - from + 1
	buffer := make([]byte, bufferSize)

	_, err := file.Seek(from, 0)
	if err != nil {
		return "", err
	}

	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer[:bufferSize]), nil
}

func blook(pattern string, file *os.File, size int64) (int64, error) {
	return -1, nil
}

//file is a file in which search will performed
//pattern is a string which should be found in file
//returns index of the first byte of the first line
//which starts with pattern
func filterLines(pattern string, file File) (int64, error) {
	if file.size == 0 {
		return -1, nil
	}

	return blook(pattern, file.ptr, file.size)
}
