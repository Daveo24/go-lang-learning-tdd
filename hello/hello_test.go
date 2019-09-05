package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, Dave"
	actual := Hello("Dave")

	if actual != expected {
		t.Errorf("actual %q expected %q", actual, expected)
	}
}
