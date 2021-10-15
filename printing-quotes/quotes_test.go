package main

import "testing"

func TestQuotes(t *testing.T) {
	expected := `Obi-Wan Kenobi says, "These aren't the droids you're looking for."`
	actual := Quotes("Obi-Wan Kenobi", "These aren't the droids you're looking for.")
	if expected != actual {
		t.Errorf("Expected %s does not match actual %s", expected, actual)
	}
}