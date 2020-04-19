package main

import "testing"

func TestX(t *testing.T) {
	// test setup
	stuff := 1
	t.Cleanup(func() {
		stuff = 0
	})

	// reminder of the test
	t.Log("stuff: ", stuff)
}
