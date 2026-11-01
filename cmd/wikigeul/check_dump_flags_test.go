//ff:func feature=cli type=command control=sequence
//ff:what TestCheckDumpFlags — a dump file and a positive -sort-mem are needed
package main

import "testing"

func TestCheckDumpFlags(t *testing.T) {
	if err := checkDumpFlags(config{SortMem: 1, DumpFile: "d.json"}); err != nil {
		t.Fatal(err)
	}
	if err := checkDumpFlags(config{SortMem: 1}); err == nil {
		t.Fatal("no dump: want an error")
	}
	if err := checkDumpFlags(config{DumpFile: "d.json"}); err == nil {
		t.Fatal("no sort memory: want an error")
	}
}
