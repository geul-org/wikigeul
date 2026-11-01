//ff:func feature=source type=helper control=selection
//ff:what writeJSONB — one decoded JSON value printed as PostgreSQL prints jsonb
package source

import (
	"encoding/json"
	"strconv"
	"strings"
)

// writeJSONB appends v (decoded with UseNumber: nil, bool, json.Number,
// string, []any, map[string]any) to b as PostgreSQL's jsonb output does
// (JsonbToCString): null, true/false, numbers as numeric prints them
// (numericText), strings escaped (writeJSONBString), arrays and objects with
// ", " and ": " (writeJSONBArray, writeJSONBObject).
func writeJSONB(b *strings.Builder, v any) {
	switch t := v.(type) {
	case nil:
		b.WriteString("null")
	case bool:
		b.WriteString(strconv.FormatBool(t))
	case json.Number:
		b.WriteString(numericText(string(t)))
	case string:
		writeJSONBString(b, t)
	case []any:
		writeJSONBArray(b, t)
	case map[string]any:
		writeJSONBObject(b, t)
	}
}
