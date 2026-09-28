package scanner

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
		Name:    "Pascal",
		Matches: []string{".pas"},
	},
	{
		Name:    "Fortran",
		Matches: []string{".f90", ".ftn"},
	},
	{
		Name:    "C",
		Matches: []string{".c", ".h"},
	},
	{
		Name:    "C++",
		Matches: []string{".cpp", ".cc", ".cxx", ".c++", ".hpp", ".hxx"},
	},
	{
		Name:    "Objective-C",
		Matches: []string{".m"},
	},
	{
		Name:    "Objective-C++",
		Matches: []string{".mm"},
	},
	{
		Name:    "C#",
		Matches: []string{".cs"},
	},
	{
		Name:    "C3",
		Matches: []string{".c3"},
	},
	{
		Name:    "D",
		Matches: []string{".d"},
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
		Name:    "Dart",
		Matches: []string{".dart"},
	},
	{
		Name:    "Swift",
		Matches: []string{".swift"},
	},
	{
		Name:    "Haskell",
		Matches: []string{".hs"},
	},
	{
		Name:    "PHP",
		Matches: []string{".php"},
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
