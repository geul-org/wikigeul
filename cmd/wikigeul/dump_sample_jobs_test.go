//ff:func feature=cli type=command control=sequence
//ff:what TestDumpSampleJobs — the first -samples decodable Q entities in number order, or the list's first Q (an undecodable one without a line); missing runs are an error
package main

import (
	"testing"

	"github.com/geul-org/wikigeul/internal/dumpfile"
)

func TestDumpSampleJobs(t *testing.T) {
	c, st := fixtureSortedDump(t, nil)
	got, err := dumpSampleJobs(c, st, nil)
	if err != nil || len(got) != 2 || got[0].Num != 1 || got[1].Num != 42 {
		t.Fatalf("%+v %v", got, err)
	}
	got, err = dumpSampleJobs(c, st, []packJob{{Kind: 'P', Num: 31}, {Kind: 'Q', Num: 7}, {Kind: 'Q', Num: 42}, {Kind: 'Q', Num: 50}})
	if err != nil || len(got) != 2 || got[0].Num != 7 || got[0].Line != "" || got[1].Num != 42 || got[1].Line == "" {
		t.Fatalf("list %+v %v", got, err)
	}
	if got, err := dumpSampleJobs(c, st, []packJob{{Kind: 'P', Num: 31}}); err != nil || got == nil || len(got) != 0 {
		t.Fatalf("no Q in list %+v %v", got, err)
	}
	if _, err := dumpSampleJobs(c, dumpfile.State{Runs: []string{"none"}}, nil); err == nil {
		t.Fatal("missing run: want an error")
	}
}
