package main

import (
	"fmt"
	"os"
	"testing"
)

var maxBufferSize = 2

func assertEqual(t *testing.T, err error, expected interface{}, actual interface{}, message string) {
	if err != nil {
		t.Error(message, err)
		return
	}
	if expected == actual {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", expected, actual)
	} else {
		message = fmt.Sprint(message, " should be: ", expected, " actual: ", actual)
	}
	t.Error(message)
}

func TestFindBorderIn1byteFile(t *testing.T) {
	file, err := os.Open("test_files/1byte.txt")
	if err != nil {
		t.Error("Can't open 1byte.txt")
	}
	defer file.Close()

	left, err := findBorder(file, 0, 0, -1, maxBufferSize)
	assertEqual(t, err, int64(-1), left, "left border from the beginning 1byte.txt")

	right, err := findBorder(file, 0, 0, 1, maxBufferSize)
	assertEqual(t, err, int64(-1), right, "right border from the end of 1byte.txt")
}

func TestFindBorderInborderFile(t *testing.T) {
	file, err := os.Open("test_files/border.txt")
	if err != nil {
		t.Error("Can't open border.txt")
	}
	defer file.Close()

	left, err := findBorder(file, 0, 32, -1, maxBufferSize)
	assertEqual(t, err, int64(22), left, "left border from the end of border.txt")

	right, err := findBorder(file, 0, 33, 1, maxBufferSize)
	assertEqual(t, err, int64(10), right, "right border from the beginning of border.txt")
}
