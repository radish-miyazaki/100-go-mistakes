package main

import "testing"

func TestLongRunning(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	// ...
}
