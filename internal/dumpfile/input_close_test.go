//ff:func feature=dumpfile type=loader control=iteration dimension=1
//ff:what TestInputClose — a running decompressor is stopped and the file closed; gzip and plain inputs close too
package dumpfile

import (
	"strings"
	"testing"
)

func TestInputClose(t *testing.T) {
	body := "[\n" + strings.Repeat("{\"id\":\"Q1\"},\n", 20000) + "]\n"
	for _, name := range []string{"d.json.bz2", "d.json.gz", "d.json"} {
		c := fixtureSortConfig(t, name, body)
		in, err := openInput(c, State{})
		if err != nil {
			t.Fatal(err)
		}
		in.r.Read(make([]byte, 10))
		in.close()
		if in.cmd != nil {
			t.Fatalf("%s: decompressor left", name)
		}
		if _, err := in.f.Stat(); err == nil {
			t.Fatalf("%s: file open", name)
		}
	}
}
