//ff:func feature=dumpfile type=loader control=iteration dimension=1
//ff:what TestTrimLine — the entity JSON without comma and newline; brackets and blank lines are no entity
package dumpfile

import "testing"

func TestTrimLine(t *testing.T) {
	if b, ok := trimLine([]byte("{\"id\":\"Q1\"},\r\n")); !ok || string(b) != `{"id":"Q1"}` {
		t.Fatalf("got %q %v", b, ok)
	}
	for _, l := range []string{"[\n", "]\n", "\n", ""} {
		if _, ok := trimLine([]byte(l)); ok {
			t.Errorf("%q: want no entity", l)
		}
	}
}
