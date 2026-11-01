//ff:func feature=memguard type=helper control=iteration dimension=1
//ff:what TestParseSize — plain numbers, K M G T with or without iB, and what is no size
package memguard

import "testing"

func TestParseSize(t *testing.T) {
	for in, want := range map[string]int64{"0": 0, "123": 123, "1K": 1 << 10, "2KiB": 2 << 10, "3M": 3 << 20, "512MiB": 512 << 20, "6GiB": 6 << 30, "1T": 1 << 40, "7B": 7, " 8 GiB ": 8 << 30} {
		if got, err := ParseSize(in); err != nil || got != want {
			t.Fatalf("%q: %d %v", in, got, err)
		}
	}
	for _, bad := range []string{"", "lots", "-1", "G"} {
		if _, err := ParseSize(bad); err == nil {
			t.Fatalf("%q: want an error", bad)
		}
	}
}
