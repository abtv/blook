package main

import (
	"os"
	"testing"
)

func TestFindBorderIn1byteFile(t *testing.T) {
	file, err := os.Open("test_files/tiny.txt")
	if err != nil {
		t.Error("Can't open tiny.txt")
	}
	defer file.Close()

	left, err := findBorder(file, 0, 0, -1, maxBufferSize)
	assertEqual(t, err, int64(-1), left, "left border from the beginning tiny.txt")

	right, err := findBorder(file, 0, 0, 1, maxBufferSize)
	assertEqual(t, err, int64(-1), right, "right border from the end of tiny.txt")
}

func TestFindBorderInborderFile(t *testing.T) {
	file, err := os.Open("test_files/small.txt")
	if err != nil {
		t.Error("Can't open small.txt")
	}
	defer file.Close()

	for i := 0; i <= 100; i++ {
		left, err := findBorder(file, 0, 32, -1, maxBufferSize)
		assertEqual(t, err, int64(22), left, "left border from the end of small.txt")

		right, err := findBorder(file, 0, 33, 1, maxBufferSize)
		assertEqual(t, err, int64(10), right, "right border from the beginning of small.txt")

		left, err = findBorder(file, 0, 16, -1, maxBufferSize)
		assertEqual(t, err, int64(10), left, "left border from the end of small.txt")

		right, err = findBorder(file, 16, 33, 1, maxBufferSize)
		assertEqual(t, err, int64(22), right, "right border from the beginning of small.txt")
	}
}
