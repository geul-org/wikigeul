//ff:func feature=pipeline type=command control=sequence
//ff:what TestOpenPackStream — a new stream; the same run resumes at its checkpoint (RECORDED-AT from it when not given); another run, no stream, a checkpoint that is not progress and failing samples are refused
package pipeline

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

func TestOpenPackStream(t *testing.T) {
	load := fixtureLoader(t)
	none := func() ([]PackJob, error) { return []PackJob{{Kind: 'Q', Num: 1}}, nil }
	c := packTestConfig(t)
	set := filepath.Join(c.Root, "set")
	s, p, err := OpenPackStream(set, c, packTestProv(c), load, none, "")
	if err != nil || p.RecordedAt != c.RecordedAt || p.Commit != c.Commit {
		t.Fatalf("%v", err)
	}
	p.Kind, p.Num, p.Docs = 'Q', 3, 2
	meta, _ := json.Marshal(p)
	s.Checkpoint(meta)
	c.Resume = true
	if _, got, err := OpenPackStream(set, c, packTestProv(c), load, none, ""); err != nil || got.Num != 3 || got.Docs != 2 {
		t.Fatalf("resume: %v", err)
	}
	c0 := c
	c0.RecordedAt = 0
	if _, got, err := OpenPackStream(set, c0, packTestProv(c0), load, none, ""); err != nil || got.RecordedAt != c.RecordedAt {
		t.Fatalf("resume without -recorded-at: %v", err)
	}
	other := c
	other.Commit = strings.Repeat("b", 40)
	if _, _, err := OpenPackStream(set, other, packTestProv(other), load, none, ""); err == nil {
		t.Fatal("other run: want error")
	}
	if _, _, err := OpenPackStream(filepath.Join(c.Root, "none"), c, packTestProv(c), load, none, ""); err == nil {
		t.Fatal("no stream: want error")
	}
	s.Checkpoint([]byte("not json"))
	if _, _, err := OpenPackStream(set, c, packTestProv(c), load, none, ""); err == nil {
		t.Fatal("not progress: want error")
	}
	bad := packTestConfig(t)
	if _, _, err := OpenPackStream(filepath.Join(bad.Root, "set"), bad, packTestProv(bad), failLoader, none, ""); err == nil {
		t.Fatal("failing samples: want error")
	}
}
