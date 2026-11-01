//ff:func feature=memguard type=helper control=sequence
//ff:what TestBindFlags — defaults; percentages and sizes read; -mem-off; a bad size and a bad percentage fail
package memguard

import (
	"flag"
	"io"
	"testing"
	"time"
)

func TestBindFlags(t *testing.T) {
	parse := func(args ...string) (Limits, error) {
		fs := flag.NewFlagSet("t", flag.ContinueOnError)
		fs.SetOutput(io.Discard)
		var l Limits
		check := BindFlags(fs, &l)
		if err := fs.Parse(args); err != nil {
			return l, err
		}
		return l, check()
	}
	l, err := parse()
	if err != nil || l.MinFree != DefaultMinFree || l.DangerFree != 0 || l.Max != 0 || l.Interval != time.Second || l.Off || l.Meminfo != procMeminfo {
		t.Fatalf("%+v %v", l, err)
	}
	l, err = parse("-mem-min-free", "40", "-mem-danger-free", "10", "-mem-max", "2GiB", "-mem-interval", "5ms", "-mem-off", "-mem-meminfo", "f")
	if err != nil || l.MinFree != .4 || l.DangerFree != .1 || l.Max != 2<<30 || l.Interval != 5*time.Millisecond || !l.Off || l.Meminfo != "f" {
		t.Fatalf("%+v %v", l, err)
	}
	if _, err := parse("-mem-max", "lots"); err == nil {
		t.Fatal("bad size")
	}
	if _, err := parse("-mem-min-free", "0"); err == nil {
		t.Fatal("no floor")
	}
}
