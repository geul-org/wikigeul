//ff:func feature=cli type=command control=sequence
//ff:what checkDumpFlags — check the flags of a run reading a dump directly and default its workers
package main

import (
	"errors"
	"flag"
)

// checkDumpFlags checks c when -dump is given: pack mode only (directory
// mode reads entities by ID at random), a positive -sort-mem. Without
// -workers it takes defaultDumpWorkers. Without -dump nothing changes.
func checkDumpFlags(c *config, fs *flag.FlagSet) error {
	if c.DumpFile == "" {
		return nil
	}
	if c.Out != "pack" {
		return errors.New("wikigeul: -dump needs -out pack")
	}
	if c.SortMem < 1 {
		return errors.New("wikigeul: -sort-mem must be at least 1 (MiB)")
	}
	workers := false
	fs.Visit(func(f *flag.Flag) { workers = workers || f.Name == "workers" })
	if !workers {
		c.Workers = defaultDumpWorkers()
	}
	return nil
}
