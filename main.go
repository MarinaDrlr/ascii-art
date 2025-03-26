package main

import (
	"ascii-art/funcs"
	"fmt"
	"os"
)

func main() {
	// Get user input (only one argument is expected)
	input, err := funcs.GetInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	// Normalize the input to match the banner map keys
	normalizedInput := funcs.NormalizeInput(input)

	// Exit silently if input is empty after normalization
	if normalizedInput == "" {
		return
	}

	// Load the standard banner font file
	banner, err := funcs.LoadBanner()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	// Generate ASCII art from the input
	asciiArt, err := funcs.GenerateASCIIArt(normalizedInput, banner)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	// Print a blank line if input was only newlines
	if len(asciiArt) == 0 {
		fmt.Println()
		return
	}

	// Display the ASCII art to the terminal
	funcs.DisplayASCIIArt(asciiArt)
}
