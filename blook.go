package main

import (
	"errors"
	"os"
	"strings"
)

// min returns minimum of two int64 numbers
func min(a int64, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

// writeBytes writes [start; stop] bytes from fromFile to toFile
func writeBytes(fromFile *os.File, start int64, stop int64, toFile *os.File, maxBufferSize int64) (int64, error) {
	var bytesWritten int64
	bytesWritten = 0
	if start > stop {
		return bytesWritten, nil
	}

	fromFile.Seek(start, 0)
	buffer := make([]byte, min(stop-start+1, maxBufferSize))
	for current := start; current < stop; {
		bufferSize := min(stop-current+1, maxBufferSize)
		if bufferSize < maxBufferSize {
			buffer = make([]byte, bufferSize)
		}

		n, err := fromFile.Read(buffer)
		if err != nil {
			return bytesWritten, err
		} else if int64(n) < bufferSize {
			return bytesWritten, errors.New("Error: unexpected end of input")
		}
		n, err = toFile.Write(buffer)
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
func newLineIndex(buffer []byte, diff int64) int {
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
		idx = idx + int(diff)
		n--
	}
}

// findBorder searches for newline symbol in [from; to]
// when diff = 1 makes forward search (`from` -> `to`)
// when diff = -1 makes backward search (`to` -> `from`)
func findBorder(file *os.File, from int64, to int64, diff int64, maxBufferSize int64) (int64, error) {
	size := to - from + int64(1)
	currentSize := min(size, maxBufferSize)

	position := from
	if diff == -1 {
		position = to - currentSize + int64(1)
	}
	buffer := make([]byte, currentSize)

	for {
		if size == 0 {
			return -1, nil
		}
		if int64(len(buffer)) != currentSize {
			buffer = make([]byte, currentSize)
		}

		file.Seek(position, 0)

		n, err := file.Read(buffer)
		if err != nil {
			return -1, err
		} else if int64(n) < currentSize {
			return -1, errors.New("Error: unexpected end of input")
		}

		idx := newLineIndex(buffer, diff)
		if idx >= 0 {
			return position + int64(idx), nil
		}

		position = position + diff*currentSize
		size = size - currentSize
		currentSize = min(size, maxBufferSize)
	}
}

// findString searches string borders
// returns (leftBorder, rightBorder, error)
func findString(file *os.File, from int64, to int64) (int64, int64, error) {
	maxBufferSize := int64(64 * 1024)
	middle := (from + to) / 2
	strFrom, err := findBorder(file, from, middle, -1, maxBufferSize)
	if err != nil {
		return -1, -1, err
	} else if strFrom == -1 {
		//no newline found, just return from position
		strFrom = from
	} else {
		//new line found, need to increment position to omit newline byte
		strFrom++
	}
	strTo, err := findBorder(file, middle+1, to, 1, maxBufferSize)
	if err != nil {
		return -1, -1, err
	} else if strTo == -1 {
		//no newline found, just return from position
		strTo = to
	} else {
		//new line found, need to decrement position to omit newline byte
		strTo--
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

// eligibleString returns true
// for forward = true if pattern >= value (or pattern is prefix for value)
// for forward = false if pattern <= value (or pattern is prefix for value)
// otherwise it returns false
func eligibleString(pattern string, value string, forward bool) bool {
	order := -1
	if forward {
		order = 1
	}
	return strings.Compare(value, pattern) == order || strings.HasPrefix(value, pattern)
}

// blook returns first byte number in the ordered `file` where `pattern` is occured as a prefix string
func blook(pattern string, file *os.File, size int64, forward bool) (int64, error) {
	if size == 0 {
		return -1, nil
	}

	result := int64(-1)
	from := int64(0)
	to := size - 1

	const maxCalls = 64
	currCall := 0

	for {
		if from < 0 || from > to || to >= size {
			return result, nil
		}

		if currCall > maxCalls {
			return -1, errors.New("MAX_CALLS_EXCEEDED")
		}

		strFrom, strTo, err := findString(file, from, to)
		if err != nil {
			return -1, err
		}
		value, err := getString(file, strFrom, strTo)
		if err != nil {
			return -1, err
		}

		if eligibleString(pattern, value, forward) {
			//it's already result, but we need to search for more results
			if forward {
				result = strFrom
				to = strFrom - int64(1)
			} else {
				result = strTo
				from = strTo + int64(2) //next byte is \n, so we need to move to the bytes after \n
			}
		} else {
			//it's not a result, we need to search for more results
			if forward {
				from = strTo + int64(1)
			} else {
				to = strFrom - int64(1)
			}
		}

		currCall++
	}
}
