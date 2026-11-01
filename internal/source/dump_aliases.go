//ff:func feature=source type=loader control=iteration dimension=2
//ff:what dumpAliases — the aliases of a dump line by language code in byte order, then dump order
package source

import (
	"maps"
	"slices"
)

// dumpAliases lists m (language → aliases) as the PG query reads alias:
// languages by the key's bytes, each language's aliases in array order
// (ordinal).
func dumpAliases(m map[string][]dumpTerm) []Text {
	var out []Text
	for _, lang := range slices.Sorted(maps.Keys(m)) {
		for _, t := range m[lang] {
			out = append(out, Text{Lang: lang, Value: t.Value})
		}
	}
	return out
}
