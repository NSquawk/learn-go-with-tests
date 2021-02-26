package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// create a buffer
	buffer := bytes.Buffer{}

	// pass buffer to Greet and write to it
	Greet(&buffer, "Chris")

	// then get string that was written to buffer
	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}
