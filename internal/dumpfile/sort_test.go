//ff:func feature=dumpfile type=command control=iteration dimension=1
//ff:what TestSort — a plain and a gzip dump sort into key order with the filter applied, the SHA-256 checked, a finished sort returned as it is
package dumpfile

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/klauspost/compress/gzip"
)

func TestSort(t *testing.T) {
	var gz bytes.Buffer
	zw := gzip.NewWriter(&gz)
	io.WriteString(zw, testDump)
	zw.Close()
	for name, body := range map[string][]byte{"wikidata-20261019-all.json": []byte(testDump), "wikidata-20261019-all.json.gz": gz.Bytes()} {
		sortCase(t, name, body)
	}
	c := fixtureSortConfig(t, "d.json", "[\n]\n")
	file := filepath.Join(t.TempDir(), "file")
	os.WriteFile(file, nil, 0o644)
	if _, err := Sort(SortConfig{Dir: filepath.Join(file, "sort")}); err == nil {
		t.Fatal("directory under a file: want an error")
	}
	os.MkdirAll(c.Dir, 0o755)
	os.WriteFile(filepath.Join(c.Dir, "state.json"), []byte("{"), 0o644)
	if _, err := Sort(c); err == nil {
		t.Fatal("bad state: want an error")
	}
	os.Remove(filepath.Join(c.Dir, "state.json"))
	c.Info.Path += ".none"
	if _, err := Sort(c); err == nil {
		t.Fatal("missing dump: want an error")
	}
}
