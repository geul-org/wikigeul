//ff:func feature=dumpfile type=writer control=sequence
//ff:what TestRunWriterPut — records put are read back in order; a put after close is an error
package dumpfile

import (
	"path/filepath"
	"testing"
)

func TestRunWriterPut(t *testing.T) {
	p := filepath.Join(t.TempDir(), "r")
	w, _ := newRunWriter(p, 1)
	if err := w.put(Record{Kind: 'Q', Num: 1, Line: 2, Data: []byte("a")}); err != nil {
		t.Fatal(err)
	}
	w.close()
	rr, ok, err := openRun(p)
	if err != nil || !ok || rr.cur.Num != 1 || rr.cur.Line != 2 || string(rr.cur.Data) != "a" {
		t.Fatalf("%v %v", ok, err)
	}
	rr.close()
	if err := w.put(Record{Kind: 'Q'}); err == nil {
		t.Fatal("after close: want an error")
	}
}
