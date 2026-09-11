package main

import (
	"errors"
	"slices"
)

type LanguageDefinition struct {
	// Name of the language (e.g. "Go", "Typescript", "C#", etc.).
	Name string
	// File extensions used by the language (with leading dot, e.g.
	// {".c", ".h"}, {".f", ".f90", ".f95"}, etc.) and/or specific files
	// (e.g. {"Makefile"}, {"CMakeLists.txt"}, etc.).
	Matches []string
}

// This is an array of the language definitions which will be used to match
// files of the project towards total metrics.
//
// TODO(everyone): add more language definitions.
var SUPPORTED_LANGUAGES = []LanguageDefinition{
	{
		Name:    "C",
		Matches: []string{".c", ".h"},
	},
	{
		Name:    "C++",
		Matches: []string{".cpp", ".cc", ".cxx", ".c++", ".hpp", ".hxx"},
	},
	{
		Name:    "C3",
		Matches: []string{".c3"},
	},
	{
		Name:    "Rust",
		Matches: []string{".rs"},
	},
	{
		Name:    "Java",
		Matches: []string{".java"},
	},
	{
		Name:    "Kotlin",
		Matches: []string{".kt"},
	},
	{
		Name:    "Go",
		Matches: []string{".go"},
	},
	{
		Name:    "JavaScript",
		Matches: []string{".js", ".cjs", ".mjs"},
	},
	{
		Name:    "TypeScript",
		Matches: []string{".ts", ".cts", ".mts"},
	},
	{
		Name:    "Ruby",
		Matches: []string{".rb"},
	},
	{
		Name:    "Python",
		Matches: []string{".py"},
	},
	{
		Name:    "R",
		Matches: []string{".R", ".r"},
	},
	{
		Name:    "Lua",
		Matches: []string{".lua"},
	},
	{
		Name:    "CSS",
		Matches: []string{".css"},
	},
	{
		Name:    "SCSS/Sass",
		Matches: []string{".scss", ".sass"},
	},
	{
		Name:    "HTML",
		Matches: []string{".html"},
	},
	{
		Name:    "Makefile",
		Matches: []string{"Makefile"},
	},
	{
		Name:    "CMake",
		Matches: []string{".cmake", "CMakeLists.txt"},
	},
}

// Source language is not supported (yet). See [SUPPORTED_LANGUAGES].
var ErrNotSupported = errors.New("source language is not supported")

// Checks if source file supported (`match` should be an extension or specific
// file, see [SourceLanguageDefinition] for more information). Also returns its
// definition (if supported).
func IsSupported(match string) (LanguageDefinition, bool) {
	for _, lang := range SUPPORTED_LANGUAGES {
		if slices.Contains(lang.Matches, match) {
			return lang, true
		}
	}

	return LanguageDefinition{}, false
}
