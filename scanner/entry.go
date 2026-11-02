package scanner

import (
	"bufio"
	"fmt"
	"io"
	"lan/languages"
	"os"
	"sync"
	"unicode"
)

type Entry struct {
	languages.LanguageDefinition

	mutex sync.Mutex

	// Total amount of the source files.
	Files int

	// Total amount of the lines of code in source files.
	Lines int

	// Total amount of non-whitespace characters.
	Weight int
}

// Scans specified file.
func (e *Entry) Scan(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open %q: %s", path, err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	// new line counter
	nlc := 0
	// characters counter
	chc := 0

	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return fmt.Errorf("error reading character: %s", err)
			}
		}

		if !unicode.IsSpace(char) {
			chc++

			if char == '\n' {
				nlc++
			}
		}
	}

	e.mutex.Lock()

	e.Files += 1
	e.Lines += nlc
	e.Weight += chc

	e.mutex.Unlock()

	return nil
}

func EntryFromDefinition(ld languages.LanguageDefinition) *Entry {
	return &Entry{
		LanguageDefinition: ld,
	}
}
