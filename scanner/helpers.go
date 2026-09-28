package scanner

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// Count amount of characters (excluding whitespaces) in file.
func countCharsInFile(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("failed to open a file: %s", err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	charsCount := 0

	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return charsCount, fmt.Errorf("error reading character: %s", err)
			}
		}

		if !unicode.IsSpace(char) {
			charsCount++
		}
	}

	return charsCount, nil
}
