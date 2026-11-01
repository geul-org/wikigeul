//ff:func feature=memguard type=helper control=sequence
//ff:what TestBump — a hit adds one, a miss resets
package memguard

import "testing"

func TestBump(t *testing.T) {
	if bump(2, true) != 3 || bump(2, false) != 0 {
		t.Fatal("bump")
	}
}
