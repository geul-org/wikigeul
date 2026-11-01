//ff:func feature=source type=loader control=sequence
//ff:what TestDumpSnaks — snaks in the dump's property order as qualifiers of one statement; a bad value is an error
package source

import (
	"encoding/json"
	"testing"
)

func TestDumpSnaks(t *testing.T) {
	m := map[string][]dumpSnak{"P2": {{Property: "P2", SnakType: "novalue"}}, "P1": {{Property: "P1", SnakType: "somevalue"}}}
	qs, err := dumpSnaks(7, []string{"P2"}, m)
	if err != nil || len(qs) != 2 || qs[0].Property != "P2" || qs[1].Property != "P1" || qs[1].StatementID != 7 {
		t.Fatalf("%+v %v", qs, err)
	}
	if qs, err := dumpSnaks(7, nil, nil); qs != nil || err != nil {
		t.Fatal("none")
	}
	m["P3"] = []dumpSnak{{DataValue: &dumpDataValue{Value: json.RawMessage(`{`)}}}
	if _, err := dumpSnaks(7, nil, m); err == nil {
		t.Fatal("want an error")
	}
}
