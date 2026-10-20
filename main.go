package main

import (
	"fmt"
	"io/fs"
	"lan/flags"
	"lan/languages"
	"lan/lib"
	"lan/scanner"
	"log"
	"path/filepath"
	"time"

	"golang.org/x/sync/errgroup"
)

func scanDir(root string) ([]string, error) {
	files := []string{}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			abs, err := filepath.Abs(path)
			if err != nil {
				return err
			}

			files = append(files, abs)
		}

		return nil
	})

	return files, err
}

func isIgnored(path string, ignored []string) (bool, error) {
	for _, pathIgnored := range ignored {
		is, err := lib.IsSubPath(pathIgnored, path)
		if err != nil {
			return false, err
		}

		if is {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	fl, err := flags.Init()
	if err != nil {
		log.Fatalf("Error parsing program flags: %s\n", err)
	}

	// start timer
	startupTime := time.Now()

	lds, err := languages.GetDefinitions()
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	fmt.Printf("Scanning %s...\n", fl.Root)

	files, err := scanDir(fl.Root)
	if err != nil {
		log.Fatalf("Error scanning directory: %s\n", err)
	}

	fmt.Printf("Analysing files...\n")

	sc := scanner.NewScanner(lds)
	errg := new(errgroup.Group)

	for _, path := range files {
		// check if path goes into the exluded directory
		ignore, err := isIgnored(path, fl.Ignore)
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}
		if ignore {
			continue
		}

		errg.Go(func() error {
			return sc.Scan(path)
		})
	}

	err = errg.Wait()
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	totalWeight := sc.TotalWeight()
	sc.SortEntries()

	fmt.Printf("\n")

	for _, e := range sc.Entries {
		if e.Weight > 0 {
			percentage := float64(e.Weight) / float64(totalWeight) * 100.0
			fmt.Printf("%.1f%%\t%d\t%s\n", percentage, e.Weight, e.Name)
		}
	}

	elapsed := time.Since(startupTime)
	fmt.Printf("\nExecution time: %fs\n", elapsed.Seconds())
}
