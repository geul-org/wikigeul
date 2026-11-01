//ff:func feature=memguard type=helper control=sequence
//ff:what TestPressureTighter — each slack the lowest of the two, the name of the one closer to caution, whichever side it is on
package memguard

import "testing"

func TestPressureTighter(t *testing.T) {
	a := pressure{Name: "a", Caution: 10, Danger: 1, Recover: 20}
	b := pressure{Name: "b", Caution: 5, Danger: 2, Recover: 30}
	want := pressure{Name: "b", Caution: 5, Danger: 1, Recover: 20}
	if a.tighter(b) != want || b.tighter(a) != want {
		t.Fatalf("%+v %+v", a.tighter(b), b.tighter(a))
	}
}
