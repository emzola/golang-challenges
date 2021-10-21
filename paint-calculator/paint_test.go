package main

import "testing"

func TestRectangularRoom(t *testing.T) {
	expected := float64(350)
	actual := SqFtRect(25, 14)
	if expected != actual {
		t.Errorf("Expected %f does not match actual %f", expected, actual)
	}
}

func TestRoundRoom(t *testing.T) {
	expected := 17671.46
	actual := SqFtCircle(150)
	if expected != actual {
		t.Errorf("Expected %f does not match actual %f", expected, actual)
	}
}

func TestLShapedRoom(t *testing.T) {
	expected := float64(36)
	actual := SqFtLShaped(5, 6, 2, 3)
	if expected != actual {
		t.Errorf("Expected %f does not match actual %f", expected, actual)
	}
}

func TestCalculator(t *testing.T) {
	expected := int16(2)
	actual := PaintCalculator(360)
	if expected != actual {
		t.Errorf("Expected %d does not match actual %d", expected, actual)
	}
}