//ff:func feature=dumpfile type=writer control=sequence
//ff:what TestSorterSpill — held records become a sorted run named in the saved state with the plain file's hash state; nothing held writes no run; an unwritable directory is an error
package dumpfile

import (
	"crypto/sha256"
	"io"
	"testing"
)

func TestSorterSpill(t *testing.T) {
	dir := t.TempDir()
	st := &State{}
	s := &sorter{c: SortConfig{Dir: dir, Threads: 1, Log: io.Discard}, st: st, in: &input{plain: true, h: sha256.New()},
		recs: []Record{{Kind: 'Q', Num: 2}, {Kind: 'P', Num: 9}}, held: 5, since: 7}
	if err := s.spill(); err != nil {
		t.Fatal(err)
	}
	if len(st.Runs) != 1 || st.Next != 1 || len(st.Hash) == 0 || s.recs != nil || s.held != 0 || s.since != 0 {
		t.Fatalf("%+v", st)
	}
	m, _ := OpenMerger(Paths(dir, st.Runs))
	r, _, _ := m.Next()
	m.Close()
	if r.Kind != 'P' {
		t.Fatal("not sorted")
	}
	if saved, _, _ := loadState(dir); saved.Next != 1 {
		t.Fatal("not saved")
	}
	s.in.plain = false
	if err := s.spill(); err != nil || len(st.Runs) != 1 {
		t.Fatalf("empty spill: %v", err)
	}
	s.c.Dir = dir + "/none"
	s.recs = []Record{{Kind: 'Q'}}
	if err := s.spill(); err == nil {
		t.Fatal("want an error")
	}
	s.recs = nil
	if err := s.spill(); err == nil {
		t.Fatal("state not saved: want an error")
	}
}
