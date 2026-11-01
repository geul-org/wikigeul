//ff:func feature=cli type=command control=sequence
//ff:what TestRun — bad flags fail before any work; the command converts through runDump, whose failure it returns
package main

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	if err := run(nil); err == nil {
		t.Fatal("no flags: want an error")
	}
	a := strings.Repeat("a", 40)
	if err := run([]string{"-commit", a, "-lib-commit", a, "-dump", "no-such-dump.json"}); err == nil {
		t.Fatal("missing dump: want an error")
	}
}
