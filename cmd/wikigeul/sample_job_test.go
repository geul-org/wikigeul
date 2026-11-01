//ff:func feature=cli type=helper control=sequence
//ff:what TestSampleJob — a decodable line or no line stays; an undecodable listed line loses its line; an undecodable sorted line is not counted
package main

import "testing"

func TestSampleJob(t *testing.T) {
	good := packJob{Kind: 'Q', Num: 1, Dump: true, Line: `{"id":"Q1"}`}
	if j, ok := sampleJob(good, false); !ok || j != good {
		t.Fatal("good")
	}
	if j, ok := sampleJob(packJob{Kind: 'Q', Num: 2, Dump: true}, true); !ok || j.Line != "" {
		t.Fatal("missing")
	}
	bad := packJob{Kind: 'Q', Num: 7, Dump: true, Line: `{"id":"Q7","labels":[]}`}
	if j, ok := sampleJob(bad, true); !ok || j.Line != "" || j.Num != 7 {
		t.Fatal("listed bad")
	}
	if _, ok := sampleJob(bad, false); ok {
		t.Fatal("sorted bad")
	}
}
