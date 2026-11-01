//ff:func feature=cli type=helper control=sequence
//ff:what TestKeepEntity — every property and the Q range, or exactly the listed entities; never another kind
package main

import (
	"testing"

	"github.com/geul-org/wikigeul/pkg/pipeline"
)

func TestKeepEntity(t *testing.T) {
	r := keepEntity(config{Config: pipeline.Config{From: 10, To: 20}}, nil)
	if !r('P', 1) || !r('Q', 10) || !r('Q', 20) || r('Q', 9) || r('Q', 21) || r('L', 10) || r('P', 1<<56) {
		t.Fatal("range")
	}
	l := keepEntity(config{Config: pipeline.Config{From: 10, To: 20}}, []pipeline.PackJob{{Kind: 'Q', Num: 5}, {Kind: 'P', Num: 31}})
	if !l('Q', 5) || !l('P', 31) || l('Q', 10) || l('P', 1) {
		t.Fatal("list")
	}
}
