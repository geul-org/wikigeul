//ff:func feature=memguard type=helper control=sequence
//ff:what TestLowerLimit — the lower of two, 0 being no limit on either side
package memguard

import "testing"

func TestLowerLimit(t *testing.T) {
	if lowerLimit(5, 3) != 3 || lowerLimit(3, 5) != 3 || lowerLimit(0, 5) != 5 || lowerLimit(5, 0) != 5 || lowerLimit(0, 0) != 0 {
		t.Fatal("lowerLimit")
	}
}
