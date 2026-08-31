package main

import (
	"fmt"
	"lan/languages"
	"os"
	"path/filepath"
)

// If source file is not supported, ScanFile will return
// [languages.ErrNotSupported].
func ScanFile(fpath string, l *languages.Languages) error {
	fbase := filepath.Base(fpath)
	fext := filepath.Ext(fpath)

	// check match for whole file
	isSupported, lang := languages.IsSupported(fbase)
	if !isSupported {
		// if not, check for file extension
		isSupported, lang = languages.IsSupported(fext)
		if !isSupported {
			return languages.ErrNotSupported
		}
	}

	lHas := l.GetLanguageByName(lang.Name) != nil
	if !lHas {
		l.AddLanguage(languages.NewLanguageStats(lang))
	}

	stats := l.GetLanguageByName(lang.Name)
	if stats == nil {
		return fmt.Errorf("cannot retrieve language from collection (nil-pointer)")
	}

	err := stats.ScanFile(fpath)
	if err != nil {
		return err
	}

	return nil
}

func ScanDirectory(path string, l *languages.Languages) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		entryPath := filepath.Join(path, entry.Name())

		if entry.IsDir() {
			ScanDirectory(entryPath, l)
		} else {
			err = ScanFile(entryPath, l)
			if err != nil {
				if err == languages.ErrNotSupported {
					continue
				}

				return fmt.Errorf("error scanning \"%s\": %s", entryPath, err)
			}
		}
	}

	return nil
}
