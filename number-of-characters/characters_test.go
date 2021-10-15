package main

import "testing"

func TestLength(t *testing.T) {
	expected := len("Homer")
	actual := Length("Homer")
	if expected != actual {
		t.Errorf("Expected %d does not match actual %d", expected, actual)
	}
}