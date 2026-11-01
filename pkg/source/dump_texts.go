//ff:func feature=source type=loader control=iteration dimension=1
//ff:what dumpTexts — the labels or descriptions of a dump line by language code in byte order
package source

import (
	"maps"
	"slices"
)

// dumpTexts lists m (language → term) as the PG query reads label and
// description: one Text per map key (wikipgeul stores the key as the language),
// ordered by the key's bytes (COLLATE "C"), the text exact (NUL included, as
// ExactValue restores it from value_json).
func dumpTexts(m map[string]dumpTerm) []Text {
	var out []Text
	for _, lang := range slices.Sorted(maps.Keys(m)) {
		out = append(out, Text{Lang: lang, Value: m[lang].Value})
	}
	return out
}
