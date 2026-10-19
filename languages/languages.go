package languages

import (
	_ "embed"
	"encoding/json"
)

//go:embed languages.json
var languagesjson []byte

type LanguageDefinition struct {
	// Name of the language definition.
	Name string `json:"name"`
	// Shell file name/extension pattern used by the language (used by the
	// [filepath.Match]).
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

// Returns length of the longest [LanguageDefinition.Name].
func LongestName(lds []LanguageDefinition) int {
	l := 0

	for _, ld := range lds {
		ldl := len(ld.Name)
		if ldl > l {
			l = ldl
		}
	}

	return l
}
