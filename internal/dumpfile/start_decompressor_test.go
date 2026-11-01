//ff:func feature=dumpfile type=loader control=iteration dimension=1
//ff:what TestStartDecompressor — gzip in-process, bzip2 through the program (with lbzip2's thread flag when it is the program); bad gzip or a missing program is an error
package dumpfile

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStartDecompressor(t *testing.T) {
	const body = "[\n]\n"
	for _, name := range []string{"d.json.gz", "d.json.bz2"} {
		c := fixtureSortConfig(t, name, body)
		f, _ := os.Open(c.Info.Path)
		in := &input{f: f, raw: f}
		if err := startDecompressor(in, c); err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		b, _ := io.ReadAll(in.r)
		in.close()
		if string(b) != body {
			t.Fatalf("%s: %q", name, b)
		}
	}
	in := &input{raw: strings.NewReader("not gzip")}
	if err := startDecompressor(in, SortConfig{Info: Info{Name: "x.gz"}}); err == nil {
		t.Fatal("bad gzip: want an error")
	}
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "lbzip2"), []byte("#!/bin/sh\necho \"$@\"\n"), 0o755)
	t.Setenv("PATH", dir)
	in = &input{raw: strings.NewReader("")}
	if err := startDecompressor(in, SortConfig{Info: Info{Name: "x.bz2"}, Threads: 3}); err != nil {
		t.Fatal(err)
	}
	b, _ := io.ReadAll(in.r)
	in.cmd.Wait()
	if string(b) != "-dc -n 3\n" {
		t.Fatalf("lbzip2 args %q", b)
	}
	t.Setenv("PATH", t.TempDir())
	in = &input{raw: strings.NewReader("")}
	if err := startDecompressor(in, SortConfig{Info: Info{Name: "x.bz2"}}); err == nil {
		t.Fatal("no bzip2: want an error")
	}
}
