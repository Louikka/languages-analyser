package lib

import (
	"path/filepath"
	"strings"
)

func IsSubPath(parent, child string) (bool, error) {
	p, err := filepath.Abs(parent)
	if err != nil {
		return false, err
	}

	c, err := filepath.Abs(child)
	if err != nil {
		return false, err
	}

	rel, err := filepath.Rel(p, c)
	if err != nil {
		return false, err
	}

	// if rel starts with ".", then there is either no difference between
	// paths (".") or child path is not a subpath of the parent ("..",
	// "../..", etc.)
	return !strings.HasPrefix(rel, "."), nil
}
