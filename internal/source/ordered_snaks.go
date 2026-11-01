//ff:func feature=source type=loader control=iteration dimension=2
//ff:what orderedSnaks — the snaks of a property map in the dump's property order, then the properties it misses, sorted (wikipg's ordinal order)
package source

import (
	"slices"
	"sort"
)

// orderedSnaks lays out m (by property) exactly as wikipg numbers the
// ordinal of qualifier and reference_snak rows (wikipg internal/load
// orderedSnaks): the properties of order (qualifiers-order, snaks-order) as
// listed, then the properties of m missing from order in sorted order, each
// property's snaks in array order. A property listed twice in order is laid
// out twice, as wikipg does.
func orderedSnaks(order []string, m map[string][]dumpSnak) []dumpSnak {
	var rest []string
	for p := range m {
		if !slices.Contains(order, p) {
			rest = append(rest, p)
		}
	}
	sort.Strings(rest)
	var out []dumpSnak
	for _, p := range append(slices.Clone(order), rest...) {
		out = append(out, m[p]...)
	}
	return out
}
