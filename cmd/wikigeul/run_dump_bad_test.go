//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestRunDumpBad — a missing dump, memory that cannot be read, a missing registry and a missing ID list stop the run before it sorts
package main

import (
	"path/filepath"
	"testing"

	"github.com/geul-org/wikigeul/pkg/memguard"
)

func TestRunDumpBad(t *testing.T) {
	dump := fixtureDumpFile(t, testDumpBody)
	reg := t.TempDir()
	fixtureRegistry(t, reg)
	none := filepath.Join(t.TempDir(), "none")
	for name, change := range map[string]func(*config){
		"dump":     func(c *config) { c.DumpFile = none },
		"memory":   func(c *config) { c.Mem = memguard.Limits{Meminfo: none} },
		"registry": func(c *config) { c.LangRoot = none },
		"ids":      func(c *config) { c.IDs = none },
	} {
		c := packTestConfig(t)
		c.DumpFile, c.LangRoot, c.SortDir, c.Mem = dump, reg, t.TempDir(), memguard.Limits{Off: true}
		change(&c)
		if err := runDump(c); err == nil {
			t.Fatalf("%s: want an error", name)
		}
	}
}
