package main

import "testing"

func TestHello(t *testing.T) {
	name := "Eric"
	got := Hello(name)
	want := "Hello, Eric!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
