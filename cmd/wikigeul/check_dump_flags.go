//ff:func feature=cli type=command control=sequence
//ff:what checkDumpFlags — check the flags of the dump file: -dump given, a positive -sort-mem
package main

import "errors"

// checkDumpFlags checks c: wikigeul reads a dump file, so -dump is required
// (the loaded database is read by wikipg export-geul), and -sort-mem must be
// positive.
func checkDumpFlags(c config) error {
	if c.DumpFile == "" {
		return errors.New("wikigeul: -dump is required (a dump loaded into PostgreSQL converts with wikipg export-geul)")
	}
	if c.SortMem < 1 {
		return errors.New("wikigeul: -sort-mem must be at least 1 (MiB)")
	}
	return nil
}
