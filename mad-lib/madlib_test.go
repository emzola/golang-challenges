package main

import "testing"

func TestOutput(t *testing.T) {
	expected := "Do you walk your blue dog quickly? That's hilarious!"
	actual := Output("dog", "walk", "blue", "quickly")
	if expected != actual {
		t.Errorf("Expected %s does not match actual %s", expected, actual)
	}
}