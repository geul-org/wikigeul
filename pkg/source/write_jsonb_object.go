//ff:func feature=source type=helper control=iteration dimension=1
//ff:what writeJSONBObject — a JSON object printed as jsonb: keys in jsonb order, ": " and ", "
package source

import "strings"

// writeJSONBObject appends m to b as {"k1": v1, "k2": v2, …} with the keys
// in jsonb's order (jsonbKeys). m already holds the last value of a key given
// twice, as jsonb keeps.
func writeJSONBObject(b *strings.Builder, m map[string]any) {
	b.WriteByte('{')
	for i, k := range jsonbKeys(m) {
		if i > 0 {
			b.WriteString(", ")
		}
		writeJSONBString(b, k)
		b.WriteString(": ")
		writeJSONB(b, m[k])
	}
	b.WriteByte('}')
}
