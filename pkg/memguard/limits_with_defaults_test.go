//ff:func feature=memguard type=helper control=sequence
//ff:what TestLimitsWithDefaults — zero settings take their defaults, times following the interval; set ones stay
package memguard

import (
	"testing"
	"time"
)

func TestLimitsWithDefaults(t *testing.T) {
	l := Limits{}.withDefaults()
	if l.MinFree != DefaultMinFree || l.DangerFree != DefaultMinFree/2 || l.Gap != defaultGap || l.Interval != time.Second ||
		l.Lead != 10*time.Second || l.Hold != 30*time.Second || l.Cooldown != 5*time.Second || l.Confirm != defaultConfirm {
		t.Fatalf("%+v", l)
	}
	s := Limits{MinFree: .4, DangerFree: .1, Gap: .2, Interval: time.Millisecond, Lead: 1, Hold: 2, Cooldown: 3, Confirm: 4}
	if got := s.withDefaults(); got.MinFree != .4 || got.DangerFree != .1 || got.Gap != .2 || got.Interval != time.Millisecond ||
		got.Lead != 1 || got.Hold != 2 || got.Cooldown != 3 || got.Confirm != 4 {
		t.Fatalf("%+v", got)
	}
	if got := (Limits{Interval: time.Millisecond}).withDefaults(); got.Hold != 30*time.Millisecond {
		t.Fatalf("%+v", got)
	}
}
