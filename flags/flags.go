package flags

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"
)

type ProgramFlags struct {
	// Absolute path to the directory where to perform analysis.
	Root string
	// List of the absolute paths which should be ignored while performing
	// analysis.
	Ignore []string
}

func InitPropgramFlags() (ProgramFlags, error) {
	ignore := flag.StringP(
		"ignore",
		"i",
		"",
		"Comma-separated list of directories that should be ignored while performing analysis.",
	)

	flag.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"Usage: %s [root_directory] [options]\n\n"+
				"Root directory is optional (defaults to the current directory).\n\n"+
				"Options:\n",
			os.Args[0],
		)
		flag.PrintDefaults()
	}

	flag.Parse()

	flags := ProgramFlags{}

	/// root

	// defaults to the current directory
	root := "."

	posArgs := flag.Args()
	if len(posArgs) > 0 {
		root = posArgs[0]
	}

	rootAbs, err := filepath.Abs(root)
	if err != nil {
		return flags, fmt.Errorf("failed to convert root directory to absolute path: %s", err)
	}

	rootAbsStats, err := os.Stat(rootAbs)
	if err != nil {
		return flags, fmt.Errorf("cannot get stats of the path %q: %s", rootAbs, err)
	}

	if !rootAbsStats.IsDir() {
		return flags, fmt.Errorf("provided path %q is not a directory", root)
	}

	flags.Root = rootAbs

	/// ignore

	split := strings.SplitSeq(*ignore, ",")
	for s := range split {
		if len(s) > 0 {
			flags.Ignore = append(flags.Ignore, filepath.Join(flags.Root, s))
		}
	}

	return flags, nil
}
