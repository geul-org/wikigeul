//ff:func feature=source type=loader control=sequence
//ff:what TestDumpQualifier — a snak in qualifier form with its statement; a value that is not JSON is an error
package source

import (
	"encoding/json"
	"testing"
)

func TestDumpQualifier(t *testing.T) {
	q, err := dumpQualifier(3, dumpSnak{SnakType: "value", Property: "P580", Datatype: "string", DataValue: &dumpDataValue{Type: "string", Value: json.RawMessage(`"v"`)}})
	if err != nil || q != (Qualifier{StatementID: 3, Property: "P580", Type: "string", Datatype: "string", Value: "v", SnakType: "value"}) {
		t.Fatalf("%+v %v", q, err)
	}
	if _, err := dumpQualifier(3, dumpSnak{DataValue: &dumpDataValue{Value: json.RawMessage(`[`)}}); err == nil {
		t.Fatal("want an error")
	}
}
