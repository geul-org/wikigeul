//ff:func feature=pipeline type=helper control=sequence
//ff:what TestTransformedBy — a document's TRANSFORMED-BY records as hex, differing by commit; bytes that are not a document give ""
package pipeline

import (
	"strings"
	"testing"
)

func TestTransformedBy(t *testing.T) {
	load := fixtureLoader(t)
	c := packTestConfig(t)
	a, _, _, _ := buildOne(load, c, packTestProv(c), PackJob{Kind: 'Q', Num: 1})
	c.Commit = strings.Repeat("b", 40)
	b, _, _, _ := buildOne(load, c, packTestProv(c), PackJob{Kind: 'Q', Num: 1})
	ta, tb := transformedBy(a), transformedBy(b)
	if ta == "" || ta == tb || !strings.HasPrefix(ta, "0x") {
		t.Fatalf("%s %s", ta, tb)
	}
	if transformedBy([]byte{1, 2}) != "" {
		t.Fatal("not a document")
	}
}
