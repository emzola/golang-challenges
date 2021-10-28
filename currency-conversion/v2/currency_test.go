package main

import "testing"

func TestConversion(t *testing.T) {
	expected := 111.38
	actual := ConvertCurrency(81, 137.51)
	if expected != actual {
		t.Errorf("Expected %f does not match actual %f", expected, actual)
	}
}