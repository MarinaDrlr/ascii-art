package funcs

import (
	"bufio"
	"fmt"
	"os"
)

// LoadBanner reads the "standard.txt" banner file and maps characters to ASCII art
func LoadBanner() (map[rune][]string, error) {
	filename := "fonts/standard.txt"

	// Check if the file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("Banner file standard.txt not found.")
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Could not open banner file standard.txt.")
	}
	defer file.Close()

	bannerMap := make(map[rune][]string)
	scanner := bufio.NewScanner(file)

	var currentChar rune = ' ' // Start from space
	var charLines []string
	linesRead := 0

	for scanner.Scan() {
		linesRead++
		line := scanner.Text()

		if line == "" {
			if len(charLines) > 0 {
				bannerMap[currentChar] = append([]string{}, charLines...)
				currentChar++
				charLines = nil
			}
			continue
		}

		charLines = append(charLines, line)
	}

	if linesRead == 0 {
		return nil, fmt.Errorf("Banner file standard.txt is empty.")
	}

	if len(charLines) > 0 {
		bannerMap[currentChar] = append([]string{}, charLines...)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Failed to read banner file standard.txt: %s", err)
	}

	// Validate map integrity: must have 95 printable ASCII characters
	if len(bannerMap) != 95 {
		return nil, fmt.Errorf("Banner file standard.txt is corrupted.")
	}
	for _, lines := range bannerMap {
		if len(lines) != 8 {
			return nil, fmt.Errorf("Banner file standard.txt is corrupted.")
		}
	}

	return bannerMap, nil
}
