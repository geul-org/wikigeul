//ff:func feature=pipeline type=command control=sequence
//ff:what TestPackWriterCheckpoint — the report is flushed and the progress saved with the stream; a report that cannot flush or sync fails
package pipeline

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestPackWriterCheckpoint(t *testing.T) {
	c := Config{Report: filepath.Join(t.TempDir(), "r.tsv"), Every: 5}
	w, _ := newPackWriter(packTestStream(t), c, PackProgress{})
	w.prog.Docs, w.since = 7, 3
	if err := w.checkpoint(); err != nil || w.since != 0 {
		t.Fatal(err)
	}
	var p PackProgress
	json.Unmarshal(w.s.Meta(), &p)
	if b, _ := os.ReadFile(c.Report); p.Docs != 7 || int64(len(b)) != p.Report {
		t.Fatalf("%+v %d", p, len(b))
	}
	w.w.WriteString("x")
	w.rep.Close()
	if err := w.checkpoint(); err == nil {
		t.Fatal("closed report: want error")
	}
	w2, _ := newPackWriter(packTestStream(t), c, PackProgress{})
	w2.w.Flush()
	w2.rep.Close()
	if err := w2.checkpoint(); err == nil {
		t.Fatal("report that cannot sync: want error")
	}
}
