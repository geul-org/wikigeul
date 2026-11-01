//ff:func feature=dumpfile type=loader control=iteration dimension=1
//ff:what TestInputFinish — after the text, the whole file's SHA-256 for plain, gzip and bzip2; a failing decompressor or an unreadable rest of the file is an error
package dumpfile

import (
	"io"
	"os"
	"testing"
)

func TestInputFinish(t *testing.T) {
	for _, name := range []string{"d.json", "d.json.gz", "d.json.bz2"} {
		c := fixtureSortConfig(t, name, "[\n]\n")
		in, err := openInput(c, State{})
		if err != nil {
			t.Fatal(err)
		}
		b, _ := io.ReadAll(in.r)
		if in.plain {
			in.h.Write(b)
		}
		sum, err := in.finish()
		in.close()
		if err != nil || sum != c.Info.SHA256 {
			t.Fatalf("%s: %s %v", name, sum, err)
		}
	}
	g := fixtureSortConfig(t, "d.json.gz", "[\n]\n")
	in, err := openInput(g, State{})
	if err != nil {
		t.Fatal(err)
	}
	io.ReadAll(in.r)
	in.f.Close()
	if _, err := in.finish(); err == nil {
		t.Fatal("file closed before the rest was hashed: want an error")
	}
	c := fixtureSortConfig(t, "d.json.bz2", "[\n]\n")
	os.WriteFile(c.Info.Path, []byte("not bzip2"), 0o644)
	in, err = openInput(c, State{})
	if err != nil {
		t.Fatal(err)
	}
	io.ReadAll(in.r)
	if _, err := in.finish(); err == nil {
		t.Fatal("bzip2 failed: want an error")
	}
	in.close()
}
