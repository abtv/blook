package main

import (
	"os"
	"testing"
)

func TestFindString(t *testing.T) {
	file, err := os.Open("test_files/border.txt")
	if err != nil {
		t.Error("Can't open border.txt")
	}
	defer file.Close()

	for i := 0; i < 100; i++ {
		left, right, err := findString(file, 0, 16)
		assertEqual(t, err, int64(0), left, "")
		assertEqual(t, err, int64(9), right, "")

		//search again the same to test work with file current position
		left, right, err = findString(file, 0, 16)
		assertEqual(t, err, int64(0), left, "")
		assertEqual(t, err, int64(9), right, "")

		left, right, err = findString(file, 0, 32)
		assertEqual(t, err, int64(11), left, "")
		assertEqual(t, err, int64(21), right, "")
	}
}
