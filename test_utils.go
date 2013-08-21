package appa

import "fmt"
import "testing"

func assertIntEquals(t *testing.T, expected int, actual int) bool {
	if actual != expected {
		t.Errorf("Expected %d, got %d.", expected, actual)
		return false
	}

	return true
}

func assertNodeStringEquals(t *testing.T, expected string, actual Node) bool {
	if actual == nil {
		t.Errorf("Expected \"%s\", got nil.", expected)
		return false
	}

	if stringer, ok := actual.(fmt.Stringer); ok {
		return assertStringerEquals(t, expected, stringer)
	}

	str := fmt.Sprint(actual)
	return assertStringEquals(t, expected, str)
}

func assertStringEquals(t *testing.T, expected string, actual string) bool {
	if actual != expected {
		t.Errorf("Expected \"%s\", got \"%s\".", expected, actual)
		return false
	}

	return true
}

func assertStringerEquals(t *testing.T, expected string, actual fmt.Stringer) bool {
	if actual == nil {
		t.Errorf("Expected \"%s\", got nil.", expected)
		return false
	}

	return assertStringEquals(t, expected, actual.String())
}
