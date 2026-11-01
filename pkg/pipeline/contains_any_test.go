//ff:func feature=pipeline type=helper control=sequence
//ff:what TestContainsAny — true when a part occurs, false when none or no parts
package pipeline

import "testing"

func TestContainsAny(t *testing.T) {
	if !containsAny("time: JSON 아님", []string{"x", "JSON"}) || containsAny("abc", []string{"d"}) || containsAny("abc", nil) {
		t.Fatal("containsAny")
	}
}
