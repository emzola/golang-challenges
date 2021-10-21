package main

import "testing"

func TestSqUnit(t *testing.T) {
	expected := int64(300)
	actual := SqUnit(15, 20)
	if expected != actual {
		t.Errorf("Expected %d does not match actual %d", expected, actual)
	}
}

func TestSqFeetToSqMeter(t *testing.T) {
	expected := 27.871
	actual := SqFeetToSqMeter(300) 
	if expected != actual {
		t.Errorf("Expected %f does not match actual %f", expected, actual)
	}
}

func TestSqMeterToSqFeet(t *testing.T) {
	expected := 3229.173
	actual := SqMeterToSqFeet(300) 
	if expected != actual {
		t.Errorf("Expected %f does not match actual %f", expected, actual)
	}
}