package funcs

import (
	"fmt"
	"os"
)

// GetInput retrieves the user's input text.
// It expects exactly 1 argument: the string to convert to ASCII art.
// If no argument is provided, or too many arguments are given, it returns an error.
func GetInput() (string, error) {
	args := os.Args[1:] // Skip the program name

	if len(args) == 0 {
		// No input provided
		return "", fmt.Errorf("No input was given.\nUsage: go run . [STRING]\nEX: go run . \"something\"")
	}

	if len(args) > 1 {
		// Too many arguments (banner selection is not supported in this version)
		return "", fmt.Errorf("Usage: go run . [STRING]\nEX: go run . \"something\"")
	}

	input := args[0] // The input string to be converted
	return input, nil
}
