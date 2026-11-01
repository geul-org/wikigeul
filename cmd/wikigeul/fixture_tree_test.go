//ff:func feature=cli type=helper control=sequence
//ff:what fixtureTree — every file under a directory with its bytes, by relative path
package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

// fixtureTree reads root for comparing two sets.
func fixtureTree(t *testing.T, root string) map[string]string {
	t.Helper()
	out := map[string]string{}
	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		b, err := os.ReadFile(p)
		rel, _ := filepath.Rel(root, p)
		out[rel] = string(b)
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	return out
}
