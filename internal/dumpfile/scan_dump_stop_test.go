//ff:func feature=dumpfile type=command control=sequence
//ff:what TestScanDumpStop — a guard stop writes what is held as a run, saves the state and returns the stop; the sort then goes on from there to the same records
package dumpfile

import (
	"errors"
	"testing"

	"github.com/geul-org/wikigeul/pkg/memguard"
)

func TestScanDumpStop(t *testing.T) {
	const body = "[\n{\"id\":\"Q3\"},\n{\"id\":\"Q1\"},\n{\"id\":\"P2\"}\n]\n"
	c := fixtureSortConfig(t, "d.json", body)
	c.Guard = fixtureGuard(t, memguard.Sample{RSS: 1, Total: 1 << 40, Avail: 1 << 36}, 1)
	var stop memguard.StopError
	if _, err := Sort(c); !errors.As(err, &stop) {
		t.Fatal(err)
	}
	st, ok, _ := loadState(c.Dir)
	if !ok || st.Scanned || st.Lines == 0 || st.Lines >= 5 {
		t.Fatalf("%+v", st)
	}
	c.Guard = nil
	done, err := Sort(c)
	if err != nil || !done.Done || done.Kept != 3 {
		t.Fatalf("%+v %v", done, err)
	}
}
