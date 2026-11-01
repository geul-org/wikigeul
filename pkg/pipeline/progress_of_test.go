//ff:func feature=pipeline type=helper control=sequence
//ff:what TestProgressOf — the run's identity, nothing done
package pipeline

import "testing"

func TestProgressOf(t *testing.T) {
	c := Config{Commit: "a", LibCommit: "b", LangDoc: "l", Schema: "s", Dump: 1, From: 2, To: 3, RecordedAt: 4}
	p := progressOf(c, "h")
	if p != (PackProgress{Commit: "a", LibCommit: "b", LangDoc: "l", Schema: "s", Dump: 1, From: 2, To: 3, IDs: "h", RecordedAt: 4}) {
		t.Fatalf("%+v", p)
	}
}
