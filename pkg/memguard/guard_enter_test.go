//ff:func feature=memguard type=helper control=sequence
//ff:what TestGuardEnter — nil lets through; no more than the active count run at once; raising the count lets a waiter in; a stop answers every waiter false
package memguard

import (
	"testing"
	"time"
)

func TestGuardEnter(t *testing.T) {
	var none *Guard
	if !none.Enter() {
		t.Fatal("nil")
	}
	none.Leave()
	g, _, _ := fixtureGuard(2)
	g.setActive(1)
	if !g.Enter() {
		t.Fatal("first")
	}
	got := make(chan bool, 2)
	wait := func() { got <- g.Enter() }
	go wait()
	go wait()
	select {
	case <-got:
		t.Fatal("second entered over the count")
	case <-time.After(20 * time.Millisecond):
	}
	g.setActive(2)
	if !<-got {
		t.Fatal("raised count")
	}
	g.apply(decision{Level: Danger, Press: pressure{Size: 1}})
	if <-got {
		t.Fatal("stop must answer false")
	}
	if g.Enter() {
		t.Fatal("after the stop")
	}
}
