//ff:func feature=pipeline type=command control=sequence
//ff:what TestBuildOne — an entity of the dump gives its document; a missing one is false without error; a failing input is an error
package pipeline

import "testing"

func TestBuildOne(t *testing.T) {
	load := fixtureLoader(t)
	c := packTestConfig(t)
	b, rep, ok, err := buildOne(load, c, packTestProv(c), PackJob{Kind: 'Q', Num: 42})
	if err != nil || !ok || len(b) == 0 || rep.ID != "Q42" || rep.Labels == 0 {
		t.Fatalf("%v %v", rep, err)
	}
	if _, rep, ok, err := buildOne(load, c, packTestProv(c), PackJob{Kind: 'Q', Num: 999}); ok || err != nil || !rep.Missing {
		t.Fatal("missing")
	}
	if _, _, ok, err := buildOne(failLoader, c, packTestProv(c), PackJob{Kind: 'Q', Num: 42}); ok || err == nil {
		t.Fatal("failing input: want error")
	}
}
