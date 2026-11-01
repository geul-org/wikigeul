//ff:func feature=memguard type=command control=sequence
//ff:what BindFlags — define the guard's flags and return the check to call after parsing
package memguard

import (
	"flag"
	"time"
)

// BindFlags defines the -mem- flags on fs, filling l. The returned function
// settles the percentages and checks l after fs.Parse.
func BindFlags(fs *flag.FlagSet, l *Limits) func() error {
	var minFree, dangerFree float64
	fs.BoolVar(&l.Off, "mem-off", false, "turn the memory guard off")
	fs.Float64Var(&minFree, "mem-min-free", DefaultMinFree*100, "percent of memory (machine or cgroup) to keep free: below it the run slows down and gives memory back")
	fs.Float64Var(&dangerFree, "mem-danger-free", 0, "percent of memory free below which the run checkpoints and stops, exit 4 (0: half of -mem-min-free)")
	fs.Func("mem-max", "cap on this process's RSS, e.g. 6GiB (default none); the run slows down near it and stops before it", func(v string) (err error) {
		l.Max, err = ParseSize(v)
		return err
	})
	fs.DurationVar(&l.Interval, "mem-interval", time.Second, "time between memory samples")
	fs.StringVar(&l.Meminfo, "mem-meminfo", procMeminfo, "file read for MemTotal and MemAvailable (tests give a file of their own)")
	return func() error {
		l.MinFree, l.DangerFree = minFree/100, dangerFree/100
		c := l.withDefaults()
		c.MinFree = l.MinFree // 0 must fail here, not become the default
		return c.check()
	}
}
