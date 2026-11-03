//ff:func feature=pipeline type=command control=sequence
//ff:what TestPackWriterMemStop — the checkpoint is saved and the StopError returned; a report that cannot be saved is its own error
package pipeline

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"

	"github.com/geul-org/wikigeul/pkg/memguard"
)

func TestPackWriterMemStop(t *testing.T) {
	c := packTestConfig(t)
	c.Report, c.Guard = filepath.Join(t.TempDir(), "r.tsv"), fixtureGuard(t, 10)
	set := filepath.Join(t.TempDir(), "set")
	o := pack.StreamOptions{Build: pack.BuildOptions{Codec: pack.Stored{}}, Issuers: []string{WikidataIssuer}, Created: 1}
	s, err := pack.NewStream(set, o)
	if err != nil {
		t.Fatal(err)
	}
	w, _ := newPackWriter(s, c, PackProgress{})
	var stop memguard.StopError
	if err := w.memStop(); !errors.As(err, &stop) {
		t.Fatal(err)
	}
	if r, err := pack.ResumeStream(set, o); err != nil || len(r.Meta()) == 0 {
		t.Fatalf("no checkpoint: %v", err)
	}
	if b, _ := os.ReadFile(c.Report); len(b) == 0 {
		t.Fatal("report not flushed")
	}
	bad, _ := newPackWriter(packTestStream(t), c, PackProgress{})
	bad.w.WriteString("x")
	bad.rep.Close()
	if err := bad.memStop(); err == nil || errors.As(err, &stop) {
		t.Fatalf("checkpoint failure: %v", err)
	}
}
