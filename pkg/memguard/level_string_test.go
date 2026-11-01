//ff:func feature=memguard type=helper control=sequence
//ff:what TestLevelString — each level's name
package memguard

import "testing"

func TestLevelString(t *testing.T) {
	if Normal.String() != "normal" || Caution.String() != "caution" || Danger.String() != "danger" {
		t.Fatal("names")
	}
}
