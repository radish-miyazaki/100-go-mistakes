package main

import (
	"strings"
	"testing"
)

func TestCountEmptyLines(t *testing.T) {
	emptyLines, err := countEmptyLines(strings.NewReader(
		`foo
					bar

					baz
					`))
	_ = emptyLines
	_ = err
}
