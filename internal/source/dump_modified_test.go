//ff:func feature=source type=loader control=sequence
//ff:what TestDumpModified — empty is 0, RFC 3339 is POSIX seconds rounded down, anything else an error
package source

import "testing"

func TestDumpModified(t *testing.T) {
	if n, err := dumpModified(""); n != 0 || err != nil {
		t.Fatal("empty")
	}
	if n, err := dumpModified("2026-10-17T04:21:25.9Z"); n != 1792210885 || err != nil {
		t.Fatalf("%d %v", n, err)
	}
	if n, err := dumpModified("1969-12-31T23:59:59.5Z"); n != -1 || err != nil {
		t.Fatalf("before 1970: %d %v", n, err)
	}
	if _, err := dumpModified("yesterday"); err == nil {
		t.Fatal("want an error")
	}
}
