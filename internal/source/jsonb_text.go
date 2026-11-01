//ff:func feature=source type=helper control=selection
//ff:what jsonbText — a JSON value as PostgreSQL gives value #>> '{}' after storing it as jsonb: a string unquoted, null empty, else the jsonb text
package source

import (
	"bytes"
	"encoding/json"
	"strings"
)

// jsonbText returns what value #>> '{}' (NULL as empty) gives for raw stored
// in a jsonb column: a top-level string unquoted, JSON null (SQL NULL from
// #>>) or no value at all empty, anything else printed as PostgreSQL prints
// jsonb (writeJSONB: keys sorted by length then bytes, the last of duplicate
// keys kept, ", " and ": " separators, numbers through numeric).
func jsonbText(raw []byte) (string, error) {
	if len(bytes.TrimSpace(raw)) == 0 {
		return "", nil
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	var v any
	if err := dec.Decode(&v); err != nil {
		return "", err
	}
	switch t := v.(type) {
	case nil:
		return "", nil
	case string:
		return t, nil
	default:
		var b strings.Builder
		writeJSONB(&b, v)
		return b.String(), nil
	}
}
