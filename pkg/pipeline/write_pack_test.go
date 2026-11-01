//ff:func feature=pipeline type=command control=sequence
//ff:what TestWritePack — the produced jobs go through the loaders into a closed set with the report; an unopenable report and a producer error fail
package pipeline

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWritePack(t *testing.T) {
	load := fixtureLoader(t)
	c := packTestConfig(t)
	prov := packTestProv(c)
	c.Report = filepath.Join(t.TempDir(), "r.tsv")
	produce := func(jobs chan<- PackJob, errp *error) {
		defer close(jobs)
		jobs <- PackJob{Seq: 0, Kind: 'Q', Num: 1}
		jobs <- PackJob{Seq: 1, Kind: 'Q', Num: 42}
		jobs <- PackJob{Seq: 2, Kind: 'Q', Num: 999}
	}
	p := PackProgress{RecordedAt: 1_800_000_000}
	if err := WritePack(packTestStream(t), c, p, prov, []Loader{load, load}, produce); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(c.Report)
	if strings.Count(string(b), "\tok\t") != 2 || strings.Count(string(b), "\tmissing\t") != 1 {
		t.Fatalf("report:\n%s", b)
	}
	c.Report = t.TempDir()
	if err := WritePack(packTestStream(t), c, p, prov, []Loader{load}, produce); err == nil {
		t.Fatal("report is a directory: want error")
	}
	c.Report = ""
	failed := errors.New("producer")
	bad := func(jobs chan<- PackJob, errp *error) {
		*errp = failed
		close(jobs)
	}
	if err := WritePack(packTestStream(t), c, p, prov, []Loader{load}, bad); !errors.Is(err, failed) {
		t.Fatalf("producer error: %v", err)
	}
}
