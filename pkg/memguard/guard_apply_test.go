//ff:func feature=memguard type=command control=sequence
//ff:what TestGuardApply — numbers published, workers set, memory given back, caution entries counted once each, the stop made once with its reason and the workers released
package memguard

import (
	"errors"
	"testing"
)

func TestGuardApply(t *testing.T) {
	g, mem, _ := fixtureGuard(8)
	p := pressure{Name: "system", Free: 1, Size: 2, Caution: 500}
	g.apply(decision{Level: Caution, Active: 4, Reclaim: true, Budget: 1 << 40, Press: p, Near: true, Reason: "why"})
	g.apply(decision{Level: Caution, Active: 4, Budget: 1 << 40, Press: p})
	if g.Cautions() != 1 || mem.reclaims != 1 || g.Active(8) != 4 || g.spare.Load() != 500 || g.budget.Load() != 1<<40 || g.near.Load() || g.reason != "why" {
		t.Fatalf("%+v", g)
	}
	g.apply(decision{Level: Normal, Active: 4, Press: p})
	g.apply(decision{Level: Caution, Active: 4, Press: p})
	if g.Cautions() != 2 {
		t.Fatal(g.Cautions())
	}
	g.SetPhase("convert")
	g.apply(decision{Level: Danger, Active: 0, Reason: "gone", Press: p})
	g.apply(decision{Level: Danger, Active: 0, Press: p})
	var stop StopError
	if !errors.As(g.Err(), &stop) || stop.Reason != "gone" || stop.Phase != "convert" {
		t.Fatalf("%+v", stop)
	}
	select {
	case <-g.Stopped():
	default:
		t.Fatal("not stopped")
	}
}
