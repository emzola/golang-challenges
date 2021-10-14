package main

import (
	"os"
	"testing"
)

func TestPromptInput(t *testing.T) {
	arg := os.Args[1]
	if arg != "Brian" {
		t.Errorf("String mismatch on test")
	}
}

func TestOutput(t *testing.T) {
	expected := "Hello, nice to meet you!"
	actual := Output()
	if expected != actual {
		t.Errorf("Expected %s does not match actual %s", expected, actual)
	}
}