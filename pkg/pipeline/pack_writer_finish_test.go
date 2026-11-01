//ff:func feature=pipeline type=command control=sequence
//ff:what TestPackWriterFinish — the stream closes into a set; a report that cannot flush and a stream that cannot close fail
package pipeline

import (
	"testing"
	"time"
)

func TestPackWriterFinish(t *testing.T) {
	load := fixtureLoader(t)
	c := packTestConfig(t)
	s := packTestStream(t)
	w, _ := newPackWriter(s, Config{Every: 10}, PackProgress{})
	w.take(convertOne(load, c, packTestProv(c), s, PackJob{Kind: 'Q', Num: 1}))
	if err := w.finish(time.Now()); err != nil {
		t.Fatal(err)
	}
	empty, _ := newPackWriter(packTestStream(t), Config{Every: 10}, PackProgress{})
	if err := empty.finish(time.Now()); err == nil {
		t.Fatal("no document: want error")
	}
	bad, _ := newPackWriter(packTestStream(t), Config{Every: 10}, PackProgress{})
	bad.w.WriteString("x")
	bad.rep.Close()
	if err := bad.finish(time.Now()); err == nil {
		t.Fatal("report that cannot flush: want error")
	}
}
