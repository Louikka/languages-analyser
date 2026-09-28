package flags

import (
	"lan/lib"
	"os"
	"slices"
	"strings"
)

// Returns all flag arguments (those which starts with "-" or "--").
func FlagArgs() []string {
	return lib.Filter(os.Args[1:], func(arg string) bool {
		return strings.HasPrefix(arg, "-")
	})
}

// Returns all positional (non-flag) arguments.
func PositionalArgs() []string {
	return lib.Filter(os.Args[1:], func(arg string) bool {
		return !strings.HasPrefix(arg, "-")
	})
}

func BoolArg(names []string) bool {
	for _, arg := range FlagArgs() {
		if slices.Contains(names, arg) {
			return true
		}
	}

	return false
}

// Returns next argument after name (or defaultValue).
func ArgWithValue(names []string, defaultValue string) string {
	args := os.Args[1:]

	for i, arg := range args {
		if slices.Contains(names, arg) && len(args) > i+1 {
			return args[i+1]
		}
	}

	return defaultValue
}
