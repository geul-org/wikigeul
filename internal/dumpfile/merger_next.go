//ff:func feature=dumpfile type=loader control=sequence
//ff:what Merger.Next — the next record of the merged runs in run order; false when every run is used up
package dumpfile

// Next returns the smallest front record and moves that run on; a run that
// ends is closed and leaves the heap.
func (m *Merger) Next() (Record, bool, error) {
	if len(m.h) == 0 {
		return Record{}, false, nil
	}
	top := m.h[0]
	r := top.cur
	ok, err := top.advance()
	if err != nil {
		return Record{}, false, err
	}
	if !ok {
		top.close()
		m.h[0] = m.h[len(m.h)-1]
		m.h = m.h[:len(m.h)-1]
	}
	siftDown(m.h, 0)
	return r, true, nil
}
