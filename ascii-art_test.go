package main

import (
	"ascii-art/funcs"
	"testing"
)

// TestNormalizeInput checks if NormalizeInput correctly replaces `\n` (literal backslash + n) with actual newlines
func TestNormalizeInput(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`Hello\nWorld`, "Hello\nWorld"},
		{`Line1\nLine2\nLine3`, "Line1\nLine2\nLine3"},
	}

	for _, test := range tests {
		result := funcs.NormalizeInput(test.input)
		if result != test.expected {
			t.Errorf("NormalizeInput(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

// TestNormalizeInputNewlines checks that actual newlines are preserved
func TestNormalizeInputNewlines(t *testing.T) {
	input := "Line1\nLine2"
	expected := "Line1\nLine2"

	result := funcs.NormalizeInput(input)
	if result != expected {
		t.Errorf("NormalizeInput with real newlines returned %q; want %q", result, expected)
	}
}

// TestLoadBanner ensures that the standard banner is loaded correctly and is not empty
func TestLoadBanner(t *testing.T) {
	banner, err := funcs.LoadBanner()
	if err != nil {
		t.Errorf("LoadBanner failed: %v", err)
	}

	if len(banner) == 0 {
		t.Errorf("Expected non-empty banner map")
	}
}

// TestGenerateASCIIArt verifies that ASCII art is generated correctly for various inputs
func TestGenerateASCIIArt(t *testing.T) {
	banner, err := funcs.LoadBanner()
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}

	tests := []struct {
		input    string
		expected int
	}{
		{"A", 8},
		{"Hello", 8},
		{"Hello\nWorld", 16},
	}

	for _, test := range tests {
		result, err := funcs.GenerateASCIIArt(test.input, banner)
		if err != nil {
			t.Errorf("GenerateASCIIArt(%q) failed: %v", test.input, err)
		}
		if len(result) != test.expected {
			t.Errorf("Expected %d lines, got %d", test.expected, len(result))
		}
	}
}

// TestGenerateASCIIArtUnsupportedChar checks error handling for unsupported characters
func TestGenerateASCIIArtUnsupportedChar(t *testing.T) {
	banner, err := funcs.LoadBanner()
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}

	_, err = funcs.GenerateASCIIArt("Hello 🌍", banner)
	if err == nil {
		t.Errorf("Expected error for unsupported character, got nil")
	}
}
