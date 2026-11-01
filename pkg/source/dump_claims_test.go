//ff:func feature=source type=loader control=sequence
//ff:what TestDumpClaims — properties as text, statements in array order, numbered from 0; a bad value is an error
package source

import (
	"encoding/json"
	"testing"
)

func TestDumpClaims(t *testing.T) {
	m := map[string][]dumpStatement{
		"P31":  {{Rank: "normal", Mainsnak: dumpSnak{Property: "P31"}}},
		"P214": {{Rank: "a", Mainsnak: dumpSnak{Property: "P214"}}, {Rank: "b", Mainsnak: dumpSnak{Property: "P214"}}},
	}
	cs, err := dumpClaims(m)
	if err != nil || len(cs) != 3 || cs[0].Rank != "a" || cs[1].Rank != "b" || cs[2].Property != "P31" || cs[2].ID != 2 {
		t.Fatalf("%+v %v", cs, err)
	}
	m["P1"] = []dumpStatement{{Mainsnak: dumpSnak{DataValue: &dumpDataValue{Value: json.RawMessage(`{`)}}}}
	if _, err := dumpClaims(m); err == nil {
		t.Fatal("want an error")
	}
}
