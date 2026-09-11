package main

import (
	"bufio"
	"cmp"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"sync"
	"unicode"

	"golang.org/x/sync/errgroup"
)

type LanguageEntry struct {
	mutex sync.Mutex

	Definition LanguageDefinition
	// Total amount of non-whitespace characters in source files of this
	// language.
	Weight int
}

func NewLanguageEntry(def LanguageDefinition) *LanguageEntry {
	return &LanguageEntry{
		Definition: def,
	}
}

// Scans source file.
func (e *LanguageEntry) Scan(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open a file: %s", err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return fmt.Errorf("error reading character: %s", err)
			}
		}

		// count everything that isn't a whitespace (or tab, newline, etc.)
		if !unicode.IsSpace(char) {
			e.mutex.Lock()
			e.Weight++
			e.mutex.Unlock()
		}
	}

	return nil
}

// scanner

type Scanner struct {
	errg errgroup.Group

	Entries []*LanguageEntry
}

func NewScanner() Scanner {
	return Scanner{}
}

// Returns pointer to the [LanguageEntry] in the list (or nil if no such entry
// exists).
func (s *Scanner) getEntry(name string) *LanguageEntry {
	for _, e := range s.Entries {
		if e.Definition.Name == name {
			return e
		}
	}

	return nil
}

// Adds new entry to the [Scanner.Entries].
func (s *Scanner) addNewEntry(entry *LanguageEntry) {
	s.Entries = append(s.Entries, entry)
}

// If source file is not supported, will return [ErrNotSupported].
func (s *Scanner) scanFile(path string) error {
	fbase := filepath.Base(path)
	fext := filepath.Ext(path)

	// check match for whole file
	lang, isSupported := IsSupported(fbase)
	if !isSupported {
		// if not, check for file extension
		lang, isSupported = IsSupported(fext)
		if !isSupported {
			return ErrNotSupported
		}
	}

	entry := s.getEntry(lang.Name)
	if entry == nil {
		entry = NewLanguageEntry(lang)
		s.addNewEntry(entry)
	}

	s.errg.Go(func() error {
		return entry.Scan(path)
	})

	return nil
}

func (s *Scanner) scanDirectory(path string) error {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range dirEntries {
		entryPath := filepath.Join(path, entry.Name())

		if entry.IsDir() {
			err := s.scanDirectory(entryPath)
			if err != nil {
				return err
			}
		} else {
			err = s.scanFile(entryPath)
			if err != nil && err != ErrNotSupported {
				return fmt.Errorf("error scanning \"%s\": %v", entryPath, err)
			}
		}
	}

	return nil
}

func (s *Scanner) ScanDirectory(path string) error {
	err := s.scanDirectory(path)
	if err != nil {
		return err
	}

	return s.errg.Wait()
}

// Sorts language entries by weight (in descending order).
func (s *Scanner) Sort() {
	slices.SortFunc(s.Entries, func(a, b *LanguageEntry) int {
		return cmp.Compare(b.Weight, a.Weight)
	})
}

func (s *Scanner) GetTotalWeight() int {
	acc := 0

	for _, e := range s.Entries {
		acc += e.Weight
	}

	return acc
}
