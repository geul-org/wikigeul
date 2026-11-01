//ff:func feature=memguard type=helper control=sequence
//ff:what TestResumeCommand — -resume added once, plain words as they are, the rest quoted for the shell
package memguard

import "testing"

func TestResumeCommand(t *testing.T) {
	if got := ResumeCommand([]string{"wikigeul", "-allow", "raw:time: 0년", "it's"}); got != `wikigeul -allow 'raw:time: 0년' 'it'\''s' -resume` {
		t.Fatal(got)
	}
	if got := ResumeCommand([]string{"wikigeul", "-resume", "-root", "/a/b"}); got != "wikigeul -resume -root /a/b" {
		t.Fatal(got)
	}
}
