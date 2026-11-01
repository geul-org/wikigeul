//ff:func feature=memguard type=helper control=sequence
//ff:what TestGuardStopped — nil channel for no guard; open until the stop, closed after
package memguard

import "testing"

func TestGuardStopped(t *testing.T) {
	var none *Guard
	if none.Stopped() != nil {
		t.Fatal("nil")
	}
	g, _, _ := fixtureGuard(1)
	select {
	case <-g.Stopped():
		t.Fatal("closed early")
	default:
	}
	g.apply(decision{Level: Danger, Press: pressure{Size: 1}})
	<-g.Stopped()
}
