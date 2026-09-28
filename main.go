package main

import (
	"fmt"
	"lan/flags"
	"lan/lib"
	"lan/scanner"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	f, err := flags.InitPropgramFlags()
	if err != nil {
		log.Fatalf("Error parsing program flags: %s\n", err)
	}

	startupTime := time.Now()

	fmt.Printf("Scanning %s...\n", f.Root)

	dirEntries, err := lib.ScanDirectory(f.Root)
	if err != nil {
		log.Fatalf("Error scanning directory: %s\n", err)
	}

	fmt.Printf("Analysing files...\n")

	errg := new(errgroup.Group)

	scanner := scanner.Init()

	for _, path := range dirEntries {
		// check if path goes into the exluded directory
		if lib.HasAnyOfPrefixes(path, f.Ignore) {
			continue
		}

		errg.Go(func() error {
			return scanner.Scan(path)
		})
	}

	err = errg.Wait()
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	totalWeight := scanner.TotalWeight()
	scanner.SortEntries()

	fmt.Print("\n")

	for _, e := range scanner.Entries {
		if e.Weight > 0 {
			percentage := float64(e.Weight) / float64(totalWeight) * 100.0
			fmt.Printf("%.1f%%\t%d\t%s\n", percentage, e.Weight, e.Name)
		}
	}

	elapsed := time.Since(startupTime)
	fmt.Printf("\nExecution time: %fs\n", elapsed.Seconds())
}
