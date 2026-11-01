//ff:func feature=cli type=helper control=sequence
//ff:what TestListedJobs — listed entities come with their dump lines, a missing one without, a repeated ID once; every entity unlisted
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/wikigeul/internal/dumpfile"
)

func TestListedJobs(t *testing.T) {
	dir := t.TempDir()
	body := "[\n{\"id\":\"Q7\"},\n{\"id\":\"P31\"},\n{\"id\":\"Q2\"},\n{\"id\":\"Q7\",\"x\":1}\n]\n"
	path := filepath.Join(dir, "wikidata-20261019-all.json")
	os.WriteFile(path, []byte(body), 0o644)
	sum := sha256.Sum256([]byte(body))
	info, _ := dumpfile.InfoOf(path, "", hex.EncodeToString(sum[:]), 0)
	st, err := dumpfile.Sort(dumpfile.SortConfig{Info: info, Dir: filepath.Join(dir, "s"), Mem: 1 << 20, Threads: 1, Log: io.Discard,
		Keep: func(byte, uint64) bool { return true }})
	if err != nil {
		t.Fatal(err)
	}
	var got []packJob
	collect := func(j packJob) bool { got = append(got, j); return true }
	m, _ := dumpfile.OpenMerger(dumpfile.Paths(filepath.Join(dir, "s"), st.Runs))
	err = dumpJobs(newDumpCursor(m), []packJob{{Kind: 'P', Num: 31}, {Kind: 'Q', Num: 3}, {Kind: 'Q', Num: 7}}, collect)
	m.Close()
	if err != nil || len(got) != 3 || got[0].Line != `{"id":"P31"}` || !got[1].Dump || got[1].Line != "" || got[2].Line != `{"id":"Q7"}` {
		t.Fatalf("listed %+v %v", got, err)
	}
	got = nil
	m, _ = dumpfile.OpenMerger(dumpfile.Paths(filepath.Join(dir, "s"), st.Runs))
	err = dumpJobs(newDumpCursor(m), nil, collect)
	m.Close()
	if err != nil || len(got) != 3 || got[0].Kind != 'P' || got[1].Num != 2 || got[2].Num != 7 {
		t.Fatalf("all %+v %v", got, err)
	}
	boom := errors.New("boom")
	if err := listedJobs(&dumpCursor{err: boom}, []packJob{{Kind: 'Q', Num: 1}}, collect); err != boom {
		t.Fatalf("cursor error: %v", err)
	}
	m, _ = dumpfile.OpenMerger(dumpfile.Paths(filepath.Join(dir, "s"), st.Runs))
	n := 0
	err = listedJobs(newDumpCursor(m), []packJob{{Kind: 'P', Num: 31}, {Kind: 'Q', Num: 2}}, func(packJob) bool { n++; return false })
	m.Close()
	if err != nil || n != 1 {
		t.Fatalf("stopped: %d %v", n, err)
	}
}
