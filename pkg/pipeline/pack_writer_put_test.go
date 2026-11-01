//ff:func feature=pipeline type=command control=sequence
//ff:what TestPackWriterPut — a document goes in as list and version, counted, checkpointing when due; a missing entity only reports; a dropped line is passed over; an error, another TRANSFORMED-BY and a key out of order fail; a strict writer stops at an anomaly before writing and passes a clean result; a failing checkpoint fails the put
package pipeline

import (
	"errors"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/wikigeul/pkg/port"
)

func TestPackWriterPut(t *testing.T) {
	load := fixtureLoader(t)
	c := packTestConfig(t)
	s := packTestStream(t)
	w, _ := newPackWriter(s, Config{Every: 1}, PackProgress{})
	conv := func(c Config, n uint64) packResult {
		return convertOne(load, c, packTestProv(c), s, PackJob{Kind: 'Q', Num: n})
	}
	if err := w.put(conv(c, 1)); err != nil || w.prog.Docs != 1 || s.Count() != 2 || string(s.Meta()) == "" {
		t.Fatalf("%v", err)
	}
	if err := w.put(packResult{Job: PackJob{Kind: 'Q', Num: 5}, Missing: true}); err != nil || w.prog.Num != 5 || w.prog.Docs != 1 {
		t.Fatal("missing")
	}
	if err := w.put(packResult{Job: PackJob{Kind: 'Q', Num: 6}, Dropped: "x"}); err != nil || w.prog.Num != 6 || w.prog.Docs != 1 {
		t.Fatal("dropped line")
	}
	other := c
	other.Commit = strings.Repeat("b", 40)
	if err := w.put(conv(other, 2)); err == nil {
		t.Fatal("other TRANSFORMED-BY: want error")
	}
	if err := w.put(conv(c, 1)); err == nil {
		t.Fatal("same document again: want error")
	}
	if err := w.put(packResult{Job: PackJob{Kind: 'Q', Num: 9}, Err: errors.New("x")}); err == nil {
		t.Fatal("failed conversion: want error")
	}
	strict, _ := newPackWriter(packTestStream(t), Config{Every: 10, Strict: true, Anomaly: filepath.Join(t.TempDir(), "a.tsv")}, PackProgress{})
	bad := conv(c, 7)
	bad.Rep.Notes = append(bad.Rep.Notes, port.Note{Kind: "skipped", Reason: "빈 값", Property: "P1"})
	if err := strict.put(bad); !errors.As(err, &AnomalyError{}) || strict.prog.Docs != 0 {
		t.Fatalf("strict anomaly: %v", err)
	}
	if err := strict.put(conv(c, 7)); err != nil || strict.prog.Docs != 1 {
		t.Fatalf("strict clean: %v", err)
	}
	lazy, _ := newPackWriter(s, Config{Every: 10}, PackProgress{})
	if err := lazy.put(conv(c, 7)); err != nil || lazy.since != 1 {
		t.Fatalf("no checkpoint yet: %v", err)
	}
	broken, _ := newPackWriter(packTestStream(t), Config{Every: 1, Report: filepath.Join(t.TempDir(), "r.tsv")}, PackProgress{})
	broken.rep.Close()
	if err := broken.put(conv(c, 7)); err == nil {
		t.Fatal("checkpoint on a closed report: want error")
	}
}
