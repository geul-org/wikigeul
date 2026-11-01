//ff:type feature=source type=model
//ff:what dumpDataValue — a snak's datavalue of a dump line: type and the value JSON as written
package source

import "encoding/json"

// dumpDataValue is a snak's datavalue: its type and the value JSON exactly
// as the dump writes it, which wikipg puts in the jsonb column value (or, with
// the escape \u0000, as text in value_json).
type dumpDataValue struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}
