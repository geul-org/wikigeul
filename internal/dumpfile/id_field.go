//ff:func feature=dumpfile type=loader control=sequence
//ff:what idField — the next member of a JSON object: its string value when the key is "id", else skipped
package dumpfile

import (
	"encoding/json"
	"fmt"
)

// idField reads one member of the object dec is in. For the key "id" it
// returns the value, which must be a string, and true; any other member's
// value is skipped.
func idField(dec *json.Decoder) (string, bool, error) {
	t, err := dec.Token()
	if err != nil {
		return "", false, err
	}
	if t != "id" {
		var skip json.RawMessage
		return "", false, dec.Decode(&skip)
	}
	var id string
	if err := dec.Decode(&id); err != nil {
		return "", false, fmt.Errorf("dump line id: %w", err)
	}
	return id, true, nil
}
