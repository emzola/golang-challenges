package main

import "testing"

func TestAddition(t *testing.T) {
	expected := uint64(4)
	actual := Add(2, 2)
	if expected != actual {
		t.Errorf("Expected %d does not match actual %d", expected, actual)
	}
}

func TestSubtraction(t *testing.T) {
	expected := uint64(0)
	actual := Subtract(2, 2)
	if expected != actual {
		t.Errorf("Expected %d does not match actual %d", expected, actual)
	}
}
func TestMultiplication(t *testing.T) {
	expected := uint64(6)
	actual := Multiply(2, 3)
	if expected != actual {
		t.Errorf("Expected %d does not match actual %d", expected, actual)
	}
}
func TestDivision(t *testing.T) {
	expected := uint64(2)
	actual := Divide(4, 2)
	if expected != actual {
		t.Errorf("Expected %d does not match actual %d", expected, actual)
	}
}