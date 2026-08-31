package main

import (
	"cmp"
	"flag"
	"fmt"
	"lan/languages"
	"log"
	"slices"
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

	l := languages.NewLanguages()

	err := ScanDirectory(flags.d, &l)
	if err != nil {
		log.Fatalf("Failed to scan directory: %s", err)
	}

	totalWeight := l.GetTotalWeight()

	slices.SortFunc(l.Collection, func(a, b languages.LanguageStats) int {
		return cmp.Compare(b.Weigth, a.Weigth)
	})

	for _, lang := range l.Collection {
		percentage := float64(lang.Weigth) / float64(totalWeight) * 100.0
		fmt.Printf("%.1f%%\t%d\t%s\n", percentage, lang.Weigth, lang.Definition.Name)
	}
}
