package languages

import (
	_ "embed"
	"encoding/json"
)

//go:embed languages.json
var languagesjson []byte

type LanguageDefinition struct {
	// Name of the language.
	Name string `json:"name"`
	// List of shell filename/extension patterns used by the language
	// (used by the [filepath.Match]).
	Matches []string `json:"matches"`
}

// Reads and parses local JSON file with languages definitions.
func GetDefinitions() ([]LanguageDefinition, error) {
	var ld []LanguageDefinition

	err := json.Unmarshal(languagesjson, &ld)
	if err != nil {
		return ld, err
	}

	return ld, nil
}
