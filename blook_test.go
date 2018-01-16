package main

import (
	"os"
	"testing"
)

func TestBlook(t *testing.T) {
	file, err := os.Open("test_files/small.txt")
	if err != nil {
		t.Error("Can't open small.txt")
	}
	defer file.Close()

	position, err := blook("aaa", file, 34)
	assertEqual(t, err, int64(0), position, "")

	position, err = blook("first", file, 34)
	assertEqual(t, err, int64(0), position, "")

	position, err = blook("second", file, 34)
	assertEqual(t, err, int64(11), position, "")

	position, err = blook("third", file, 34)
	assertEqual(t, err, int64(23), position, "")

	position, err = blook("zzz", file, 34)
	assertEqual(t, err, int64(-1), position, "")
}

func TestEmpty(t *testing.T) {
	file, err := os.Open("test_files/empty.txt")
	if err != nil {
		t.Error("Can't open empty.txt")
	}
	defer file.Close()

	position, err := blook("aaa", file, 0)
	assertEqual(t, err, int64(-1), position, "")
}
