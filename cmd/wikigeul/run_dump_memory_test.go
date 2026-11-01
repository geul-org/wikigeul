//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestRunDumpMemory — memory that runs out in each phase stops the run with a StopError; the same command with -resume gives the set and report of a run never stopped
package main

import (
	"errors"
	"maps"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/wikigeul/pkg/memguard"
)

func TestRunDumpMemory(t *testing.T) {
	dump := fixtureDumpFile(t, testDumpBody)
	reg := t.TempDir()
	fixtureRegistry(t, reg)
	a := strings.Repeat("a", 40)
	base := config{DumpFile: dump, SortMem: 1}
	base.LangRoot, base.From, base.To, base.Commit, base.LibCommit, base.Workers, base.Every, base.RecordedAt, base.Samples, base.Schema, base.Dump = reg, 1, 100, a, a, 2, 1, 1_800_000_000, 2, "wikidata", 1
	full := func(c config) (string, config) {
		c.Root = t.TempDir()
		c.SortDir, c.Report = filepath.Join(c.Root, "sort"), filepath.Join(c.Root, "r.tsv")
		return filepath.Join(c.Root, ".well-known/geul/0x00010000"), c
	}
	set, ref := full(base)
	ref.Mem.Off = true
	if err := runDump(ref); err != nil {
		t.Fatal(err)
	}
	want := fixtureTree(t, set)
	wantRep, _ := os.ReadFile(ref.Report)
	for _, phase := range []string{"sort:scan", "sort:merge", "convert", "close:data", "close:p", "close:index"} {
		set, c := full(base)
		c.Mem = memguard.Limits{Confirm: 1, Read: fixturePressedReader(phase)}
		var stop memguard.StopError
		if err := runDump(c); !errors.As(err, &stop) || stop.Phase != phase {
			t.Fatalf("%s: %v", phase, err)
		}
		c.Mem, c.Resume = memguard.Limits{Off: true}, true
		if err := runDump(c); err != nil {
			t.Fatalf("%s resumed: %v", phase, err)
		}
		rep, _ := os.ReadFile(c.Report)
		if !maps.Equal(fixtureTree(t, set), want) || string(rep) != string(wantRep) {
			t.Fatalf("%s: output differs", phase)
		}
	}
}
