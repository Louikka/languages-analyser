package lib

import (
	"os"
	"path/filepath"
)

// ScanDirectory scans provided directory recursively and returns array of
// absolute paths of the found files.
func ScanDirectory(path string) ([]string, error) {
	var files = []string{}

	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return files, err
	}

	for _, e := range dirEntries {
		eAbsPath := filepath.Join(path, e.Name())

		if e.IsDir() {
			f, err := ScanDirectory(eAbsPath)
			if err != nil {
				return files, err
			}

			files = append(files, f...)
		} else {
			files = append(files, eAbsPath)
		}
	}

	return files, nil
}
