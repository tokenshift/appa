package appa

import "testing"

func assertFloatEquals(t *testing.T, expected float64, actual float64) bool {
	if actual != expected {
		t.Errorf("Expected %f, got %f.", expected, actual)
		return false
	}
	
	return true
}

func assertIntEquals(t *testing.T, expected int, actual int) bool {
	if actual != expected {
		t.Errorf("Expected %d, got %d.", expected, actual)
		return false
	}

	return true
}

func assertNil(t *testing.T, actual interface{}) bool {
	if actual != nil {
		t.Errorf("Expected nil, got %v.", actual)
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
