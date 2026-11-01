//ff:func feature=pipeline type=helper control=sequence
//ff:what TestSameRun — progress fields are ignored, any identity field differing is another run
package pipeline

import "testing"

func TestSameRun(t *testing.T) {
	want := PackProgress{Commit: "a", RecordedAt: 1}
	if !sameRun(PackProgress{Commit: "a", RecordedAt: 1, Kind: 'Q', Num: 5, Docs: 9, Report: 100}, want) {
		t.Fatal("same run")
	}
	if sameRun(PackProgress{Commit: "b", RecordedAt: 1}, want) || sameRun(PackProgress{Commit: "a", RecordedAt: 2}, want) {
		t.Fatal("other run")
	}
}
