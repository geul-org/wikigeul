//ff:func feature=dumpfile type=loader control=iteration dimension=1
//ff:what TestSorterLine — every line counts and moves the offset; kept lines are held, others counted as skipped or not read (logged); a plain file's bytes are hashed
package dumpfile

import (
	"bytes"
	"crypto/sha256"
	"strings"
	"testing"
)

func TestSorterLine(t *testing.T) {
	var log bytes.Buffer
	st := &State{}
	s := &sorter{c: SortConfig{Keep: func(k byte, n uint64) bool { return k == 'Q' }, Log: &log}, st: st, in: &input{plain: true, h: sha256.New()}}
	for _, l := range []string{"[\n", "{\"id\":\"Q1\"},\n", "{\"id\":\"P1\"},\n", "{\"x\":1},\n"} {
		s.line([]byte(l))
	}
	if st.Lines != 4 || st.Kept != 1 || st.Skipped != 1 || st.Bad != 1 || len(s.recs) != 1 || s.held != 11 || st.Offset != 37 || s.since != 37 {
		t.Fatalf("%+v %d %d", st, len(s.recs), s.held)
	}
	if !strings.Contains(log.String(), "dump line 4 not read") {
		t.Fatalf("log %q", log.String())
	}
	if s.in.h.Size() != 32 || len(s.in.h.Sum(nil)) != 32 {
		t.Fatal("hash")
	}
	s.in.plain = false
	s.line([]byte("]\n"))
	if st.Lines != 5 {
		t.Fatal("compressed line")
	}
}
