//ff:func feature=memguard type=helper control=iteration dimension=1
//ff:what TestLimitsCheck — the defaults pass; shares out of range, danger above caution, a negative cap and no interval fail
package memguard

import "testing"

func TestLimitsCheck(t *testing.T) {
	good := Limits{}.withDefaults()
	if err := good.check(); err != nil {
		t.Fatal(err)
	}
	for _, change := range []func(*Limits){
		func(l *Limits) { l.MinFree = 0 },
		func(l *Limits) { l.MinFree = .99 },
		func(l *Limits) { l.DangerFree = -.1 },
		func(l *Limits) { l.DangerFree = .5 },
		func(l *Limits) { l.Max = -1 },
		func(l *Limits) { l.Interval = 0 },
	} {
		l := good
		change(&l)
		if l.check() == nil {
			t.Fatalf("%+v: want an error", l)
		}
	}
}
