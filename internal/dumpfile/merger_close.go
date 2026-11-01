//ff:func feature=dumpfile type=loader control=iteration dimension=1
//ff:what Merger.Close — close every run still open
package dumpfile

// Close closes the runs that have not ended.
func (m *Merger) Close() {
	for _, rr := range m.h {
		rr.close()
	}
	m.h = nil
}
