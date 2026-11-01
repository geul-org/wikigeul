//ff:func feature=source type=helper control=iteration dimension=1
//ff:what writeJSONBString — a string as a JSON literal the way PostgreSQL's escape_json writes it
package source

import (
	"fmt"
	"strings"
)

// writeJSONBString appends s quoted as PostgreSQL's escape_json does: \b \f
// \n \r \t, \" and \\ by name, other bytes below 0x20 as \u00xx (lower-case
// hex), everything else — "/", DEL, non-ASCII UTF-8 — as it is.
func writeJSONBString(b *strings.Builder, s string) {
	b.WriteByte('"')
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case c == '\b':
			b.WriteString(`\b`)
		case c == '\f':
			b.WriteString(`\f`)
		case c == '\n':
			b.WriteString(`\n`)
		case c == '\r':
			b.WriteString(`\r`)
		case c == '\t':
			b.WriteString(`\t`)
		case c == '"':
			b.WriteString(`\"`)
		case c == '\\':
			b.WriteString(`\\`)
		case c < 0x20:
			fmt.Fprintf(b, `\u%04x`, c)
		default:
			b.WriteByte(c)
		}
	}
	b.WriteByte('"')
}
