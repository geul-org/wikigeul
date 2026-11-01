//ff:func feature=cli type=helper control=sequence
//ff:what defaultDumpWorkers — conversion goroutines when reading a dump directly: 70% of the CPUs, no database backends to count
package main

import "runtime"

// defaultDumpWorkers is 70% of the CPUs, at least 1: reading a dump directly
// has no PostgreSQL backend per worker (defaultWorkers halves for those).
// The same number of threads decompresses and compresses while sorting.
func defaultDumpWorkers() int {
	return max(1, runtime.NumCPU()*7/10)
}
