//ff:func feature=dumpfile type=helper control=iteration dimension=1
//ff:what TestSameSort — equal identities match whatever the progress; any identity field differing does not
package dumpfile

import "testing"

func TestSameSort(t *testing.T) {
	a := State{Name: "n", Size: 1, ModTime: 2, SHA256: "s", Filter: "f"}
	b := a
	b.Runs, b.Offset, b.Done = []string{"r"}, 9, true
	if !sameSort(a, b) {
		t.Fatal("progress compared")
	}
	for _, c := range []State{{Name: "x", Size: 1, ModTime: 2, SHA256: "s", Filter: "f"}, {Name: "n", Size: 9, ModTime: 2, SHA256: "s", Filter: "f"},
		{Name: "n", Size: 1, ModTime: 9, SHA256: "s", Filter: "f"}, {Name: "n", Size: 1, ModTime: 2, SHA256: "x", Filter: "f"}, {Name: "n", Size: 1, ModTime: 2, SHA256: "s", Filter: "x"}} {
		if sameSort(a, c) {
			t.Errorf("%+v matched", c)
		}
	}
}
