//ff:func feature=dumpfile type=writer control=sequence
//ff:what TestRunWriterClose — closing ends the stream and the file; closing twice is an error
package dumpfile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunWriterClose(t *testing.T) {
	p := filepath.Join(t.TempDir(), "r")
	w, _ := newRunWriter(p, 1)
	w.put(Record{Kind: 'P', Num: 1, Data: []byte("x")})
	if err := w.close(); err != nil {
		t.Fatal(err)
	}
	if st, err := os.Stat(p); err != nil || st.Size() == 0 {
		t.Fatal("empty file")
	}
	if err := w.close(); err == nil {
		t.Fatal("twice: want an error")
	}
}
