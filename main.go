package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"time"
)

type ProgramFlags struct {
	// Directory where to perform analysis.
	d string
}

func initPropgramFlags() *ProgramFlags {
	d := flag.String("d", ".", "directory where to perform analysis")

	flag.Parse()

	return &ProgramFlags{
		d: *d,
	}
}

func main() {
	flags := initPropgramFlags()

	startupTime := time.Now()

	dAbs, err := filepath.Abs(flags.d)
	if err != nil {
		log.Fatalf("Failed to convert target to absolute path: %s", err)
	}

	fmt.Printf("Scanning %s...", dAbs)

	scanner := NewScanner()

	err = scanner.ScanDirectory(dAbs)
	if err != nil {
		log.Fatalf("\n\nFailed to scan directory: %s", err)
	}

	totalWeight := scanner.GetTotalWeight()
	scanner.Sort()

	fmt.Print("\n\n")

	for _, lang := range scanner.Entries {
		percentage := float64(lang.Weight) / float64(totalWeight) * 100.0
		fmt.Printf("%.1f%%\t%d\t%s\n", percentage, lang.Weight, lang.Definition.Name)
	}

	elapsed := time.Since(startupTime)
	fmt.Printf("\nExecution time: %fs\n", elapsed.Seconds())
}
