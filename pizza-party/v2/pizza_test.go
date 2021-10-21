package main

import "testing"

func TestDividePizza(t *testing.T) {
	expected := int8(2)
	actual := DividePizza(2, 8, 8)
	if expected != actual {
		t.Errorf("Expected %d does not match actual %d", expected, actual)
	}
}