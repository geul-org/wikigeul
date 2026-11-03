//ff:func feature=pipeline type=helper control=sequence
//ff:what TestRunLine — the comment names range, schema, dump and both commits
package pipeline

import "testing"

func TestRunLine(t *testing.T) {
	got := RunLine(Config{From: 1, To: 10000, Schema: "wikidata", Dump: 1, Commit: "aa", LibCommit: "bb", LangDoc: "0x00000000/codecs/0x000007a0@1 0x26fb"})
	if got != "# wikigeul -from 1 -to 10000 -schema wikidata -dump-id 1 -commit aa -lib-commit bb -langreg 0x00000000/codecs/0x000007a0@1 0x26fb" {
		t.Fatalf("got %q", got)
	}
	got = RunLine(Config{From: 1, To: 2, Source: "x.json sha256:ff", Commit: "aa", LibCommit: "bb", LangDoc: "l"})
	if got != "# wikigeul -from 1 -to 2 -dump x.json sha256:ff -commit aa -lib-commit bb -langreg l" {
		t.Fatalf("dump: %q", got)
	}
}
