//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestParseFlags — defaults keep the checkpoint identity and the dump's workers; dump flags and overrides; no -dump, a short commit, a zero -sort-mem and the flags that moved to wikipg export-geul fail
package main

import (
	"strings"
	"testing"
)

func TestParseFlags(t *testing.T) {
	a, b := strings.Repeat("a", 40), strings.Repeat("b", 40)
	need := []string{"-commit", a, "-lib-commit", b, "-dump", "d.json"}
	c, err := parseFlags(need)
	if err != nil || c.Root != "." || c.From != 1 || c.To != 1000 || c.Commit != a || c.LibCommit != b || c.Schema != "wikidata" || c.Dump != 1 || c.Workers != defaultDumpWorkers() || c.DumpFile != "d.json" || c.SortDir != "wikigeul-dumpsort" || c.SortMem != 1024 {
		t.Fatalf("defaults: %+v %v", c, err)
	}
	c, err = parseFlags(append(need, "-dump-url", "u", "-dump-sha256", "ff", "-dump-time", "7", "-sort-dir", "s", "-sort-mem", "9", "-workers", "3"))
	if err != nil || c.DumpURL != "u" || c.DumpSHA256 != "ff" || c.DumpTime != 7 || c.SortDir != "s" || c.SortMem != 9 || c.Workers != 3 {
		t.Fatalf("dump: %+v %v", c, err)
	}
	for _, args := range [][]string{
		{"-commit", a, "-lib-commit", b},
		{"-commit", "abc", "-lib-commit", b, "-dump", "d.json"},
		append(need, "-sort-mem", "0"),
		append(need, "-dsn", "x"),
		append(need, "-schema", "x"),
		append(need, "-dump-id", "2"),
		append(need, "-out", "pack"),
	} {
		if _, err := parseFlags(args); err == nil {
			t.Fatalf("%v: want error", args)
		}
	}
}
