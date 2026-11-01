//ff:func feature=pipeline type=command control=sequence
//ff:what TestBuildAll — documents of the jobs the dump has; an error stops it
package pipeline

import "testing"

func TestBuildAll(t *testing.T) {
	load := fixtureLoader(t)
	c := packTestConfig(t)
	bs, err := buildAll(load, c, packTestProv(c), []PackJob{{Kind: 'Q', Num: 1}, {Kind: 'Q', Num: 999}, {Kind: 'Q', Num: 42}})
	if err != nil || len(bs) != 2 {
		t.Fatalf("%d %v", len(bs), err)
	}
	if _, err := buildAll(failLoader, c, packTestProv(c), []PackJob{{Kind: 'Q', Num: 1}}); err == nil {
		t.Fatal("failing input: want error")
	}
}
