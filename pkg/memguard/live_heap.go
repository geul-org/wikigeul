//ff:func feature=memguard type=loader control=sequence
//ff:what liveHeap — the bytes of live heap the last collection marked
package memguard

import "runtime/metrics"

// liveHeap reads /gc/heap/live:bytes.
func liveHeap() int64 {
	s := []metrics.Sample{{Name: "/gc/heap/live:bytes"}}
	metrics.Read(s)
	return int64(s[0].Value.Uint64())
}
