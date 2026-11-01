//ff:func feature=dumpfile type=loader control=iteration dimension=1
//ff:what lineID — the "id" of a dump line's entity object, read without decoding the rest of the line
package dumpfile

import (
	"bytes"
	"encoding/json"
	"errors"
)

// lineID reads the top-level members of the JSON object b in order until
// the key "id" (idField) and returns its string value. Wikidata writes "type"
// and "id" first, so a line of megabytes is not decoded here. A line that is
// not an object or has no string "id" is an error. The full decoding later
// (source.DecodeDumpEntity) checks that it finds the same ID.
func lineID(b []byte) (string, error) {
	dec := json.NewDecoder(bytes.NewReader(b))
	if t, err := dec.Token(); err != nil || t != json.Delim('{') {
		return "", errors.New("dump line is not a JSON object")
	}
	for dec.More() {
		id, found, err := idField(dec)
		if err != nil || found {
			return id, err
		}
	}
	return "", errors.New("dump line has no id")
}
