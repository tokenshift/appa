package appa

import (
	"testing"
)

func assertIntEquals(t *testing.T, expected int, actual int) bool {
	if actual != expected {
		t.Errorf("Expected %d, got %d.", expected, actual)
		return false
	}

	return true
}

func assertStringEquals(t *testing.T, expected string, actual string) bool {
	if actual != expected {
		t.Errorf("Expected \"%s\", got \"%s\".", expected, actual)
		return false
	}

	return true
}
