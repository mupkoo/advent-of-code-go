package util

import (
	"runtime"
	"testing"
)

// not a real test... used o run copy to clipboard without a main package
func TestCopyToClipboard(t *testing.T) {
	// Run this test only on mac
	if runtime.GOOS != "darwin" {
		t.Skip("Skipping test on non-macOS")
	}

	err := CopyToClipboard("advent of code")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
