//ff:func feature=cli type=command control=sequence
//ff:what TestCheckDumpFlags — nothing without -dump; pack mode and a positive -sort-mem needed; workers default to the dump default unless given
package main

import (
	"flag"
	"testing"
)

func TestCheckDumpFlags(t *testing.T) {
	parse := func(c *config, args ...string) *flag.FlagSet {
		fs := flag.NewFlagSet("t", flag.ContinueOnError)
		fs.IntVar(&c.Workers, "workers", 3, "")
		fs.Parse(args)
		return fs
	}
	c := config{Out: "pack", SortMem: 1}
	if err := checkDumpFlags(&c, parse(&c)); err != nil || c.Workers != 3 {
		t.Fatalf("no dump: %v %d", err, c.Workers)
	}
	c = config{Out: "pack", SortMem: 1, DumpFile: "d.json"}
	if err := checkDumpFlags(&c, parse(&c)); err != nil || c.Workers != defaultDumpWorkers() {
		t.Fatalf("default workers: %v %d", err, c.Workers)
	}
	c = config{Out: "pack", SortMem: 1, DumpFile: "d.json"}
	if err := checkDumpFlags(&c, parse(&c, "-workers", "5")); err != nil || c.Workers != 5 {
		t.Fatalf("given workers: %v %d", err, c.Workers)
	}
	c = config{Out: "dir", SortMem: 1, DumpFile: "d.json"}
	if err := checkDumpFlags(&c, parse(&c)); err == nil {
		t.Fatal("dir mode: want an error")
	}
	c = config{Out: "pack", DumpFile: "d.json"}
	if err := checkDumpFlags(&c, parse(&c)); err == nil {
		t.Fatal("no sort memory: want an error")
	}
}
