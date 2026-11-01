//ff:func feature=cli type=helper control=sequence
//ff:what TestSampleJob — a decodable line or no line stays; an undecodable listed line loses its line; an undecodable sorted line is not counted
package main

import (
	"testing"

	"github.com/geul-org/wikigeul/pkg/pipeline"
)

func TestSampleJob(t *testing.T) {
	good := pipeline.PackJob{Kind: 'Q', Num: 1, Line: `{"id":"Q1"}`}
	if j, ok := sampleJob(good, false); !ok || j != good {
		t.Fatal("good")
	}
	if j, ok := sampleJob(pipeline.PackJob{Kind: 'Q', Num: 2}, true); !ok || j.Line != "" {
		t.Fatal("missing")
	}
	bad := pipeline.PackJob{Kind: 'Q', Num: 7, Line: `{"id":"Q7","labels":[]}`}
	if j, ok := sampleJob(bad, true); !ok || j.Line != "" || j.Num != 7 {
		t.Fatal("listed bad")
	}
	if _, ok := sampleJob(bad, false); ok {
		t.Fatal("sorted bad")
	}
}
