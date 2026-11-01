//ff:func feature=pipeline type=command control=sequence
//ff:what TestDictSamples — the listed entities' documents, leaving out missing ones; a failing list or input fails
package pipeline

import (
	"errors"
	"testing"
)

func TestDictSamples(t *testing.T) {
	load := fixtureLoader(t)
	c := packTestConfig(t)
	list := func() ([]PackJob, error) {
		return []PackJob{{Kind: 'Q', Num: 42}, {Kind: 'Q', Num: 999}, {Kind: 'Q', Num: 1}}, nil
	}
	if bs, err := dictSamples(load, c, packTestProv(c), list); err != nil || len(bs) != 2 {
		t.Fatalf("%d %v", len(bs), err)
	}
	if _, err := dictSamples(failLoader, c, packTestProv(c), list); err == nil {
		t.Fatal("failing input: want error")
	}
	bad := func() ([]PackJob, error) { return nil, errors.New("no list") }
	if _, err := dictSamples(load, c, packTestProv(c), bad); err == nil {
		t.Fatal("failing list: want error")
	}
}
