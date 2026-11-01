//ff:func feature=dumpfile type=loader control=iteration dimension=1
//ff:what TestOpenInput — plain, gzip and bzip2 dumps open at the saved offset; a missing file, a bad gzip or an offset past the end is an error
package dumpfile

import (
	"io"
	"os"
	"testing"
)

func TestOpenInput(t *testing.T) {
	const body = "[\n{\"id\":\"Q1\"}\n]\n"
	for _, name := range []string{"d.json", "d.json.gz", "d.json.bz2"} {
		c := fixtureSortConfig(t, name, body)
		in, err := openInput(c, State{})
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		b, _ := io.ReadAll(in.r)
		in.close()
		if string(b) != body {
			t.Fatalf("%s: %q", name, b)
		}
		in, err = openInput(c, State{Offset: 2, Hash: hashState(body[:2])})
		if err != nil {
			t.Fatalf("%s at 2: %v", name, err)
		}
		b, _ = io.ReadAll(in.r)
		in.close()
		if string(b) != body[2:] {
			t.Fatalf("%s at 2: %q", name, b)
		}
	}
	c := fixtureSortConfig(t, "d.json.gz", body)
	if _, err := openInput(c, State{Offset: 999}); err == nil {
		t.Fatal("past the end: want an error")
	}
	os.WriteFile(c.Info.Path, []byte("not gzip"), 0o644)
	if _, err := openInput(c, State{}); err == nil {
		t.Fatal("bad gzip: want an error")
	}
	c.Info.Path += ".none"
	if _, err := openInput(c, State{}); err == nil {
		t.Fatal("missing: want an error")
	}
}
