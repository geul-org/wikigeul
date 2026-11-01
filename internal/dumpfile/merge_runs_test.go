//ff:func feature=dumpfile type=writer control=iteration dimension=1
//ff:what TestMergeRuns — runs written and merged yield every record in kind, number, line order, data intact
package dumpfile

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/klauspost/compress/zstd"
)

func TestMergeRuns(t *testing.T) {
	dir := t.TempDir()
	a, b, out := filepath.Join(dir, "a"), filepath.Join(dir, "b"), filepath.Join(dir, "out")
	if err := writeRun(a, []Record{{'P', 31, 9, []byte("p31")}, {'Q', 2, 3, []byte("q2")}, {'Q', 5, 1, []byte("q5")}}, 1); err != nil {
		t.Fatal(err)
	}
	if err := writeRun(b, []Record{{'Q', 1, 7, []byte("q1")}, {'Q', 2, 8, []byte("q2 again")}}, 1); err != nil {
		t.Fatal(err)
	}
	if err := mergeRuns(out, []string{a, b}, 1); err != nil {
		t.Fatal(err)
	}
	m, err := OpenMerger([]string{out})
	if err != nil {
		t.Fatal(err)
	}
	defer m.Close()
	want := []string{"p31", "q1", "q2", "q2 again", "q5"}
	var got []string
	for r, ok, err := m.Next(); ok || err != nil; r, ok, err = m.Next() {
		if err != nil {
			t.Fatal(err)
		}
		got = append(got, string(r.Data))
	}
	if !slices.Equal(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
	if err := mergeRuns(out+"2", []string{filepath.Join(dir, "none")}, 1); err == nil {
		t.Fatal("missing input: want an error")
	}
	if err := mergeRuns(filepath.Join(dir, "no", "out"), []string{a}, 1); err == nil {
		t.Fatal("unwritable output: want an error")
	}
	z, _ := zstd.NewWriter(nil)
	bad := filepath.Join(dir, "bad")
	os.WriteFile(bad, z.EncodeAll(append(appendRecord(nil, Record{Kind: 'Q', Num: 1}), 'Q'), nil), 0o644)
	if err := mergeRuns(out+"3", []string{bad}, 1); err == nil {
		t.Fatal("corrupt input: want an error")
	}
}
