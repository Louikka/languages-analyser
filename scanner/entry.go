package scanner

import (
	"sync"
)

type Entry struct {
	LanguageDefinition

	mutex sync.Mutex

	// Total amount of non-whitespace characters in source files of this
	// language.
	Weight int
}

func EntryFromDefinition(ld LanguageDefinition) *Entry {
	return &Entry{
		LanguageDefinition: ld,
	}
}

// Locks mutex, adds weight, then unlocks mutex.
func (e *Entry) AddWeight(n int) {
	e.mutex.Lock()
	e.Weight += n
	e.mutex.Unlock()
}
