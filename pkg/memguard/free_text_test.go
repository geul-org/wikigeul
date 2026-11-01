//ff:func feature=memguard type=helper control=sequence
//ff:what TestFreeText — the free share as a percentage
package memguard

import "testing"

func TestFreeText(t *testing.T) {
	if got := freeText(pressure{Free: 123, Size: 1000}); got != "12.3%" {
		t.Fatal(got)
	}
}
