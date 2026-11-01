//ff:func feature=memguard type=command control=sequence
//ff:what TestGuardPace — nil goes on; going on sets the phase; memory run out at a new phase stops at that very safe point with Confirm 1
package memguard

import (
	"errors"
	"io"
	"testing"
)

func TestGuardPace(t *testing.T) {
	var none *Guard
	if err := none.Pace("x"); err != nil {
		t.Fatal(err)
	}
	read := func(phase string) (Sample, error) {
		if phase == "close:h" {
			return fixtureSample(5), nil
		}
		return fixtureSample(60), nil
	}
	g, err := Start(Limits{Read: read, Confirm: 1}, 2, io.Discard)
	if err != nil {
		t.Fatal(err)
	}
	defer g.Close()
	if err := g.Pace("close:p"); err != nil || g.Phase() != "close:p" {
		t.Fatalf("%v %s", err, g.Phase())
	}
	var stop StopError
	if err := g.Pace("close:h"); !errors.As(err, &stop) || stop.Phase != "close:h" {
		t.Fatal(err)
	}
}
