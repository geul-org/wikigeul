//ff:func feature=dumpfile type=loader control=sequence
//ff:what TestIDField — another member is skipped, the id member gives its string; a cut-off object or a non-string id is an error
package dumpfile

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestIDField(t *testing.T) {
	dec := json.NewDecoder(strings.NewReader(`{"type":{"a":[1]},"id":"Q1"}`))
	dec.Token()
	if id, found, err := idField(dec); id != "" || found || err != nil {
		t.Fatalf("skip: %q %v %v", id, found, err)
	}
	if id, found, err := idField(dec); id != "Q1" || !found || err != nil {
		t.Fatalf("id: %q %v %v", id, found, err)
	}
	dec = json.NewDecoder(strings.NewReader(`{"id":7}`))
	dec.Token()
	if _, _, err := idField(dec); err == nil {
		t.Fatal("number id: want an error")
	}
	dec = json.NewDecoder(strings.NewReader(`{`))
	dec.Token()
	if _, _, err := idField(dec); err == nil {
		t.Fatal("cut off: want an error")
	}
}
