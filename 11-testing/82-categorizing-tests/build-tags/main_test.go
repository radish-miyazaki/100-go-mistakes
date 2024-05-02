//go:build integration

package main

import (
	"os"
	"testing"
)

func TestInsert1(t *testing.T) {
	// ...
}

func TestInsert2(t *testing.T) {
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip("skipping integration test")
	}
}
