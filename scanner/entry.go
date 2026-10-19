package scanner

import (
	"lan/languages"
	"sync"
)

type Entry struct {
	languages.LanguageDefinition

	mutex sync.Mutex

	// Total amount of non-whitespace characters in source files of this
	// language.
	Weight int
}

// Scans specified file.
func (e *Entry) Scan(path string) error {
	count, err := countCharsInFile(path)
	if err != nil {
		return err
	}

	e.mutex.Lock()

	e.Weight += count

	e.mutex.Unlock()

	return nil
}

func EntryFromDefinition(ld languages.LanguageDefinition) *Entry {
	return &Entry{
		LanguageDefinition: ld,
	}
}
