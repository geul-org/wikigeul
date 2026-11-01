//ff:func feature=dumpfile type=writer control=sequence
//ff:what TestAppendRecord — header big-endian (kind, number, line, length) then the data, after what the buffer held
package dumpfile

import (
	"bytes"
	"testing"
)

func TestAppendRecord(t *testing.T) {
	got := appendRecord([]byte{9}, Record{Kind: 'Q', Num: 0x0102, Line: 3, Data: []byte("xy")})
	want := []byte{9, 'Q', 0, 0, 0, 0, 0, 0, 1, 2, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 2, 'x', 'y'}
	if !bytes.Equal(got, want) || len(want)-1-2 != recordHeaderSize {
		t.Fatalf("% x", got)
	}
}
