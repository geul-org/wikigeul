//ff:func feature=cli type=command control=sequence
//ff:what TestRunDumpPack — a sorted dump streams into the set; -resume with no stream yet starts one; an existing set, missing runs or an unopenable report fail
package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/wikigeul/internal/dumpfile"
)

func TestRunDumpPack(t *testing.T) {
	c, st := fixtureSortedDump(t, nil)
	reg := t.TempDir()
	fixtureRegistry(t, reg)
	rd, _ := registryDir(reg)
	c.Langs, c.LangDoc, _ = languageRegistry(rd, 0)
	prov := dumpProvenance(dumpfile.Info{URL: "u", SHA256: strings.Repeat("ab", 32)}, c.Commit, c.LibCommit)
	c.Resume, c.RecordedAt = true, 0
	if err := runDumpPack(c, prov, st, nil, ""); err != nil {
		t.Fatalf("resume with no stream: %v", err)
	}
	if _, err := os.Stat(filepath.Join(c.Root, ".well-known", "geul", "0x00010000", "latest.txt")); err != nil {
		t.Fatal("no set written")
	}
	c.Resume = false
	if err := runDumpPack(c, prov, st, nil, ""); err == nil {
		t.Fatal("set exists: want an error")
	}
	if err := runDumpPack(c, prov, dumpfile.State{Runs: []string{"none"}}, nil, ""); err == nil {
		t.Fatal("missing runs: want an error")
	}
	c.Root, c.Report = t.TempDir(), t.TempDir()
	if err := runDumpPack(c, prov, st, nil, ""); err == nil {
		t.Fatal("report is a directory: want an error")
	}
}
