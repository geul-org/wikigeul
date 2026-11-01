//ff:func feature=dumpfile type=loader control=iteration dimension=1
//ff:what TestRunReaderAdvance — records in turn, false at the end; a run cut inside a header or inside the data is an error
package dumpfile

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/klauspost/compress/zstd"
)

func TestRunReaderAdvance(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "r")
	writeRun(p, []Record{{Kind: 'Q', Num: 1, Data: []byte("a")}, {Kind: 'Q', Num: 2, Line: 7, Data: []byte("bc")}}, 1)
	rr, _, _ := openRun(p)
	if ok, err := rr.advance(); !ok || err != nil || rr.cur.Num != 2 || rr.cur.Line != 7 || string(rr.cur.Data) != "bc" {
		t.Fatalf("second: %v %v", ok, err)
	}
	if ok, err := rr.advance(); ok || err != nil {
		t.Fatalf("end: %v %v", ok, err)
	}
	rr.close()
	cut := func(name string, b []byte) string {
		z, _ := zstd.NewWriter(nil)
		q := filepath.Join(dir, name)
		os.WriteFile(q, z.EncodeAll(b, nil), 0o644)
		return q
	}
	whole := appendRecord(nil, Record{Kind: 'Q', Num: 1, Data: []byte("abc")})
	for _, q := range []string{cut("head", append(append([]byte{}, whole...), 'Q', 0)), cut("data", whole[:len(whole)-1])} {
		rr, ok, err := openRun(q)
		if ok && err == nil {
			ok, err = rr.advance()
		}
		if err == nil {
			t.Errorf("%s: want an error (%v)", q, ok)
		}
	}
}
