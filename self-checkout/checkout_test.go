package main

import "testing"

func TestSubTotal(t *testing.T) {

	testData := []Cart{
		{
			price: 2500,
			quantity: 2,
		},
		{
			price: 1000,
			quantity: 1,
		},
		{
			price: 400,
			quantity: 1,
		},
	}
	expected := 64.00
	actual := SubTotal(testData)
	if expected != actual {
		t.Errorf("Expected %f does not match actual %f", expected, actual)
	}
}
func TestTaxRate(t *testing.T) {
	expected := 3.52
	actual := TaxRate(64.00)
	if expected != actual {
		t.Errorf("Expected %f does not match actual %f", expected, actual)
	}
}

func TestTotal(t *testing.T) {
	expected := 67.52
	actual := Total(64.00, 3.52)
	if expected != actual {
		t.Errorf("Expected %f does not match actual %f", expected, actual)
	}
}