package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Dave")
	want := "Hello, Dave"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
