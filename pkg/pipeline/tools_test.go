//ff:func feature=pipeline type=helper control=sequence
//ff:what TestTools — wikigeul as document 1 and the GEUL library as document 2, at their commits
package pipeline

import "testing"

func TestTools(t *testing.T) {
	ts := Tools("a", "b")
	if len(ts) != 2 || ts[0].Doc != 1 || ts[0].Commit != "a" || ts[1].Doc != 2 || ts[1].Commit != "b" {
		t.Fatalf("%+v", ts)
	}
}
