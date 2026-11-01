//ff:func feature=cli type=command control=sequence
//ff:what runDump — convert a dump read directly: its provenance, the language registry, the sort into key order, then the pack stream
package main

import (
	"fmt"

	"github.com/geul-org/wikigeul/internal/dumpfile"
)

// runDump is run for -dump: no database. The dump's provenance comes from
// the file (dumpfile.InfoOf: URL, SHA-256, download time as wikipg records
// them), the registry is read once as in run, each reference occurrence
// keeps its own snak order (the dump has it; RefOrder), and the run is named
// by the file and its SHA-256 (c.Source, part of the checkpoint identity).
// The dump is sorted into document key order first (sortDump: Stream.Put
// needs ascending keys and a dump is in no ID order), then converted as in
// pack mode (runDumpPack).
func runDump(c config) error {
	info, err := dumpfile.InfoOf(c.DumpFile, c.DumpURL, c.DumpSHA256, c.DumpTime)
	if err != nil {
		return err
	}
	reg, err := registryDir(c.LangRoot)
	if err != nil {
		return err
	}
	if c.Langs, c.LangDoc, err = languageRegistry(reg, c.LangVer); err != nil {
		return err
	}
	c.RefOrder = true
	c.Source = info.Name + " sha256:" + info.SHA256
	fmt.Printf("language registry: %s, %d tags\n", c.LangDoc, len(c.Langs))
	fmt.Printf("dump: %s (%s), %s\n", info.Path, c.Source, info.URL)
	fmt.Printf("reference snak order: per occurrence (dump) %v\n", c.RefOrder)
	ids, idsHash, err := readIDs(c.IDs)
	if err != nil {
		return err
	}
	st, err := sortDump(c, info, ids, idsHash)
	if err != nil {
		return fmt.Errorf("dump sort in %s: %w", c.SortDir, err)
	}
	return runDumpPack(c, dumpProvenance(info, c.Commit, c.LibCommit), st, ids, idsHash)
}
