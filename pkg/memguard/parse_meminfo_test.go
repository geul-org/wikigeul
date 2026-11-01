//ff:func feature=memguard type=loader control=sequence
//ff:what TestParseMeminfo — the two lines in bytes, other and short lines ignored; either missing is an error
package memguard

import "testing"

func TestParseMeminfo(t *testing.T) {
	total, avail, err := parseMeminfo([]byte("MemTotal:  100 kB\nMemFree: 1 kB\n\nx\nMemAvailable: 0 kB\n"))
	if err != nil || total != 100<<10 || avail != 0 {
		t.Fatalf("%d %d %v", total, avail, err)
	}
	if _, _, err := parseMeminfo([]byte("MemTotal: 100 kB\n")); err == nil {
		t.Fatal("no MemAvailable")
	}
	if _, _, err := parseMeminfo([]byte("MemAvailable: 100 kB\n")); err == nil {
		t.Fatal("no MemTotal")
	}
}
