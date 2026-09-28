package lib

import "strings"

// Reports whether the string s begins with any of the prefixes.
func HasAnyOfPrefixes(s string, prefixes []string) bool {
	for _, p := range prefixes {
		if strings.HasPrefix(s, p) {
			return true
		}
	}

	return false
}

func Filter[T any](slice []T, fn func(T) bool) []T {
	var r []T

	for _, v := range slice {
		if fn(v) {
			r = append(r, v)
		}
	}

	return r
}
