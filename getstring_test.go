package main

import (
	"os"
	"testing"
)

func TestGetString(t *testing.T) {
	file, err := os.Open("test_files/border.txt")
	if err != nil {
		t.Error("Can't open border.txt")
	}
	defer file.Close()

	str, err := getString(file, 0, 0)
	assertEqual(t, err, "f", str, "")

	str, err = getString(file, 0, 4)
	assertEqual(t, err, "first", str, "")

	str, err = getString(file, 11, 21)
	assertEqual(t, err, "second line", str, "")
}
