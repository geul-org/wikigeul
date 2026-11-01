//ff:func feature=source type=helper control=iteration dimension=1
//ff:what writeJSONBArray — a JSON array printed as jsonb: elements in order, ", " between
package source

import "strings"

// writeJSONBArray appends a to b as [e1, e2, …] (writeJSONB for each).
func writeJSONBArray(b *strings.Builder, a []any) {
	b.WriteByte('[')
	for i, e := range a {
		if i > 0 {
			b.WriteString(", ")
		}
		writeJSONB(b, e)
	}
	b.WriteByte(']')
}
