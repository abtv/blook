package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteBytes(t *testing.T) {
	fromName := "test_files/small.txt"
	fromFile, err := os.Open(fromName)
	if err != nil {
		t.Error("Can't open small.txt")
	}
	defer fromFile.Close()

	tmpFile, err := ioutil.TempFile("", "output")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tmpFile.Name())

	maxBufferSize := int64(2)

	bytesWritten, err := writeBytes(fromFile, 0, 34, tmpFile, maxBufferSize)
	assertEqual(t, err, int64(34), bytesWritten, "")

	fromBytes, err := ioutil.ReadFile(fromName)
	if err != nil {
		t.Error(err)
	}

	toBytes, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(fromBytes, toBytes) {
		t.Error("toFile should be the same as fromFile")
	}
}
