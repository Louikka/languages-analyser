package scanner

import (
	"cmp"
	"lan/languages"
	"path/filepath"
	"slices"
)

type Scanner struct {
	Entries []*Entry
}

func NewScanner(lds []languages.LanguageDefinition) Scanner {
	var es = []*Entry{}

	for _, ld := range lds {
		e := EntryFromDefinition(ld)
		es = append(es, e)
	}

	return Scanner{
		Entries: es,
	}
}

// Finds language entry that matches specified filename. If no entry was
// found, returns [languages.ErrNotSupported] error.
func (s Scanner) Find(name string) (*Entry, error) {
	for _, e := range s.Entries {
		for _, match := range e.Matches {
			matched, err := filepath.Match(match, name)
			if err != nil {
				return nil, err
			}
			if matched {
				return e, nil
			}
		}
	}

	return nil, languages.ErrNotSupported
}

// Checks if source file is supported, and if it is, scans it.
func (s Scanner) Scan(path string) error {
	e, err := s.Find(filepath.Base(path))
	if err != nil {
		// skip if language is not supported
		if err == languages.ErrNotSupported {
			return nil
		}

		return err
	}

	return e.Scan(path)
}

// Sorts [Scanner.Entries] by weight (in descending order) in place.
func (s Scanner) SortEntries() {
	slices.SortFunc(s.Entries, func(a, b *Entry) int {
		return cmp.Compare(b.Weight, a.Weight)
	})
}

// Calculates total weight of the entries.
func (s Scanner) TotalWeight() int {
	acc := 0

	for _, e := range s.Entries {
		acc += e.Weight
	}

	return acc
}
