//ff:func feature=cli type=helper control=sequence
//ff:what TestDefaultDumpWorkers — 70% of the CPUs, at least one
package main

import (
	"runtime"
	"testing"
)

func TestDefaultDumpWorkers(t *testing.T) {
	if n := defaultDumpWorkers(); n < 1 || n != max(1, runtime.NumCPU()*7/10) {
		t.Fatalf("%d", n)
	}
}
