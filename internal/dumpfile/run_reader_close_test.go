//ff:func feature=dumpfile type=loader control=sequence
//ff:what TestRunReaderClose — after close the run's file is closed
package dumpfile

import (
	"path/filepath"
	"testing"
)

func TestRunReaderClose(t *testing.T) {
	p := filepath.Join(t.TempDir(), "r")
	writeRun(p, []Record{{Kind: 'Q', Num: 1}}, 1)
	rr, _, _ := openRun(p)
	rr.close()
	if _, err := rr.f.Stat(); err == nil {
		t.Fatal("file still open")
	}
}
