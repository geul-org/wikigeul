//ff:func feature=dumpfile type=loader control=sequence
//ff:what TestLoadState — no file is false; a saved state reads back; a directory as the file or bad JSON is an error
package dumpfile

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadState(t *testing.T) {
	dir := t.TempDir()
	if _, ok, err := loadState(dir); ok || err != nil {
		t.Fatal("none")
	}
	want := State{Name: "n", Runs: []string{"r"}, Hash: []byte{1}, Done: true}
	saveState(dir, want)
	if got, ok, err := loadState(dir); !ok || err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("%+v %v %v", got, ok, err)
	}
	os.WriteFile(filepath.Join(dir, "state.json"), []byte("{"), 0o644)
	if _, _, err := loadState(dir); err == nil {
		t.Fatal("bad JSON: want an error")
	}
	odd := t.TempDir()
	os.Mkdir(filepath.Join(odd, "state.json"), 0o755)
	if _, _, err := loadState(odd); err == nil {
		t.Fatal("directory: want an error")
	}
}
