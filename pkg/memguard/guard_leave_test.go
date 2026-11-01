//ff:func feature=memguard type=helper control=sequence
//ff:what TestGuardLeave — a job that ends lets the next one in
package memguard

import "testing"

func TestGuardLeave(t *testing.T) {
	g, _, _ := fixtureGuard(1)
	g.Enter()
	got := make(chan bool)
	go func() { got <- g.Enter() }()
	g.Leave()
	if !<-got {
		t.Fatal("next one")
	}
}
