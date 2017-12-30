package main

import (
	"fmt"
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
