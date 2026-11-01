//ff:func feature=source type=helper control=sequence
//ff:what jsonbKeys — an object's keys in jsonb's stored order: shorter first, then by bytes
package source

import (
	"cmp"
	"maps"
	"slices"
	"strings"
)

// jsonbKeys returns the keys of m as jsonb stores and prints them
// (lengthCompareJsonbString): by length in bytes, then by bytes (memcmp).
func jsonbKeys(m map[string]any) []string {
	return slices.SortedFunc(maps.Keys(m), func(a, b string) int {
		return cmp.Or(cmp.Compare(len(a), len(b)), strings.Compare(a, b))
	})
}
