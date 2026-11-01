//ff:func feature=cli type=command control=sequence
//ff:what runDump — convert a dump read directly: its provenance, the language registry, the sort into key order, then the pack stream
package main

import (
	"fmt"
	"os"

	"github.com/geul-org/wikigeul/internal/dumpfile"
	"github.com/geul-org/wikigeul/pkg/memguard"
	"github.com/geul-org/wikigeul/pkg/pipeline"
)

// runDump converts c.DumpFile. The dump's provenance comes from the file
// (dumpfile.InfoOf: URL, SHA-256, download time as wikipgeul records them), the
// registry is read once for the run, each reference occurrence keeps its own
// snak order (the dump has it), and the run is named by the file and its
// SHA-256 (c.Source, part of the checkpoint identity). The dump is sorted
// into document key order first (sortDump: Stream.Put needs ascending keys
// and a dump is in no ID order), then converted (runDumpPack). The memory
// guard watches both (memguard.Start).
func runDump(c config) error {
	info, err := dumpfile.InfoOf(c.DumpFile, c.DumpURL, c.DumpSHA256, c.DumpTime)
	if err != nil {
		return err
	}
	if c.Guard, err = memguard.Start(c.Mem, c.Workers, os.Stdout); err != nil {
		return err
	}
	defer c.Guard.Close()
	if err := pipeline.LoadLanguages(&c.Config); err != nil {
		return err
	}
	c.Source = info.Name + " sha256:" + info.SHA256
	fmt.Printf("language registry: %s, %d tags\n", c.LangDoc, len(c.Langs))
	fmt.Printf("dump: %s (%s), %s\n", info.Path, c.Source, info.URL)
	fmt.Printf("reference snak order: per occurrence (dump) true\n")
	ids, idsHash, err := pipeline.ReadIDs(c.IDs)
	if err != nil {
		return err
	}
	st, err := sortDump(c, info, ids, idsHash)
	if err != nil {
		return fmt.Errorf("dump sort in %s: %w", c.SortDir, err)
	}
	if err := c.Guard.Err(); err != nil {
		return err
	}
	return runDumpPack(c, dumpProvenance(info, c.Commit, c.LibCommit), st, ids, idsHash)
}
