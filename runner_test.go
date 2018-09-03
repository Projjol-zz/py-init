package main

import "testing"

func TestSanitizeInput(t *testing.T) {
	inputText := "snake_case"
	if sanitizeInput(inputText) != "snake case" {
		t.Error("expected sanitized output!")
	}
}
