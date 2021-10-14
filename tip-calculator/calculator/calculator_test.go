package calculator

import "testing"

func TestCalculator(t *testing.T) {
	expected := float64(10)
	actual := Calculate(100, 10)
	if expected != actual {
		t.Errorf("Expected %f does not match actual %f", expected, actual)
	}
}