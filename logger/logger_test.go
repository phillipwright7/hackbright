package logger

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)

	if logger == nil {
		t.Error("New() returned nil")
	} else {
		logger.Log("Hello, world!")

		if buf.String() != "Hello, world!\n" {
			t.Errorf("Expected 'Hello, world!\\n', got '%s'", buf.String())
		}
	}
}
