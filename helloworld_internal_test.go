package main

import "testing"

func TestGreet(t *testing.T) {
	expectedGreeting := "Hello world"
	greeting := greet()
	if greeting != expectedGreeting {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
	}
}
