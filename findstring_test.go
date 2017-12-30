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

	left, right, err := findString(file, 0, 16)
	assertEqual(t, err, int64(0), left, "")
	assertEqual(t, err, int64(10), right, "") //also contains \n symbol in the end of line

	left, right, err = findString(file, 0, 33)
	assertEqual(t, err, int64(11), left, "")
	assertEqual(t, err, int64(22), right, "")
}
