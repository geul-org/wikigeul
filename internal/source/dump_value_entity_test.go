//ff:func feature=source type=loader control=sequence
//ff:what TestDumpValueEntity — the id of an entity value; other types, values that are not objects or have no id give none
package source

import (
	"encoding/json"
	"testing"
)

func TestDumpValueEntity(t *testing.T) {
	ent := func(v string) *dumpDataValue {
		return &dumpDataValue{Type: "wikibase-entityid", Value: json.RawMessage(v)}
	}
	if got := dumpValueEntity(ent(`{"entity-type":"item","numeric-id":5,"id":"Q5"}`)); got != "Q5" {
		t.Fatalf("got %q", got)
	}
	if got := dumpValueEntity(ent(`{"entity-type":"item","numeric-id":5}`)); got != "" {
		t.Fatalf("no id: %q", got)
	}
	if got := dumpValueEntity(ent(`"Q5"`)); got != "" {
		t.Fatalf("not an object: %q", got)
	}
	if got := dumpValueEntity(&dumpDataValue{Type: "string", Value: json.RawMessage(`{"id":"Q5"}`)}); got != "" {
		t.Fatalf("string: %q", got)
	}
}
