package util

import "testing"

// not a real test... used o run copy to clipboard without a main package
func TestCopyToClipboard(t *testing.T) {
	err := CopyToClipboard("advent of code")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
