//ff:func feature=dumpfile type=loader control=iteration dimension=1
//ff:what TestLineID — the id member is found after others; no object, no id or a non-string id is an error
package dumpfile

import "testing"

func TestLineID(t *testing.T) {
	if id, err := lineID([]byte(`{"type":"item","labels":{"en":{"value":"x"}},"id":"Q42","claims":{}}`)); err != nil || id != "Q42" {
		t.Fatalf("got %q %v", id, err)
	}
	for _, bad := range []string{`[1]`, `{"type":"item"}`, `{"id":5}`, `{"id":`} {
		if _, err := lineID([]byte(bad)); err == nil {
			t.Errorf("%s: want an error", bad)
		}
	}
}
