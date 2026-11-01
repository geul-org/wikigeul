//ff:func feature=memguard type=loader control=sequence
//ff:what TestLiveHeap — the live heap after a collection is above 0
package memguard

import (
	"runtime"
	"testing"
)

func TestLiveHeap(t *testing.T) {
	runtime.GC()
	if liveHeap() <= 0 {
		t.Fatal(liveHeap())
	}
}
