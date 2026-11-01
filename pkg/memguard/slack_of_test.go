//ff:func feature=memguard type=helper control=sequence
//ff:what TestSlackOf — bytes above each floor, negative below it
package memguard

import "testing"

func TestSlackOf(t *testing.T) {
	p := slackOf("x", 100, 1000, .3, .1, .5)
	if p.Name != "x" || p.Free != 100 || p.Size != 1000 || p.Caution != -200 || p.Danger != 0 || p.Recover != -400 {
		t.Fatalf("%+v", p)
	}
}
