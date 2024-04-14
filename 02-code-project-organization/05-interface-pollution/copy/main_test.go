package copy

import (
	"bytes"
	"strings"
	"testing"
)

func TestCopySourceToDest(t *testing.T) {
	const input = "foo"
	source := strings.NewReader(input)
	dest := bytes.NewBuffer(make([]byte, 0))

	if err := copySourceToDest(source, dest); err != nil {
		t.FailNow()
	}

	got := dest.String()
	if got != input {
		t.Errorf("expected %q, got %q", input, got)
	}
}
