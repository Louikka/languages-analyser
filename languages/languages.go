package languages

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type LanguageStats struct {
	Definition LanguageDefinition

	// Total amount of non-whitespace characters in source files of this
	// language.
	Weigth int
}

func NewLanguageStats(def LanguageDefinition) LanguageStats {
	return LanguageStats{
		Definition: def,
		Weigth:     0,
	}
}

func (stats *LanguageStats) ScanFile(fpath string) error {
	f, err := os.Open(fpath)
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
			stats.Weigth++
		}
	}

	return nil
}

type Languages struct {
	Collection []LanguageStats
}

func NewLanguages() Languages {
	return Languages{}
}

// Returns pointer to the [LanguageStats] in collection (or nil if no language
// exists in collection).
func (l *Languages) GetLanguageByName(name string) *LanguageStats {
	for i := range l.Collection {
		lang := &l.Collection[i]
		if lang.Definition.Name == name {
			return lang
		}
	}

	return nil
}

func (l Languages) GetTotalWeight() int {
	acc := 0

	for _, lang := range l.Collection {
		acc += lang.Weigth
	}

	return acc
}

// Adds new language (or, if the language already exists, merges language
// weight to the existing one).
func (l *Languages) AddLanguage(lang LanguageStats) {
	existing := l.GetLanguageByName(lang.Definition.Name)
	if existing != nil {
		existing.Weigth += lang.Weigth
	} else {
		l.Collection = append(l.Collection, lang)
	}
}
