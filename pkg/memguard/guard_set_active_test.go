//ff:func feature=memguard type=command control=sequence
//ff:what TestGuardSetActive — a change is heard once, an unchanged count not at all, 0 not at all
package memguard

import "testing"

func TestGuardSetActive(t *testing.T) {
	g, _, _ := fixtureGuard(8)
	var heard []int
	g.OnActive(func(n int) { heard = append(heard, n) })
	g.setActive(8)
	g.setActive(4)
	g.setActive(4)
	g.setActive(0)
	if len(heard) != 1 || heard[0] != 4 {
		t.Fatalf("%v", heard)
	}
	g.OnActive(nil)
	g.setActive(2)
	if g.Active(8) != 2 {
		t.Fatal(g.Active(8))
	}
}
