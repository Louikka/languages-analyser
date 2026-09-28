package scanner

import (
	"cmp"
	"slices"
	"strings"
)

type Scanner struct {
	Entries []*Entry
}

func Init() Scanner {
	var es = []*Entry{}

	for _, ld := range SUPPORTED_LANGUAGES {
		e := EntryFromDefinition(ld)
		es = append(es, e)
	}

	return Scanner{
		Entries: es,
	}
}

// Finds language entry from path.
func (s Scanner) Find(path string) (*Entry, bool) {
	for _, e := range s.Entries {
		for _, match := range e.Matches {
			if strings.HasSuffix(path, match) {
				return e, true
			}
		}
	}

	return nil, false
}

// Checks if source file is supported, and if it is, scans specified file and
// adds its weight to the entry.
func (s Scanner) Scan(path string) error {
	e, ok := s.Find(path)
	if ok {
		count, err := countCharsInFile(path)
		if err != nil {
			return err
		}

		e.AddWeight(count)
	}

	return nil
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
