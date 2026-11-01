//ff:func feature=dumpfile type=loader control=iteration dimension=1
//ff:what OpenMerger — open sorted runs as one merged stream of records
package dumpfile

// OpenMerger opens every run of paths (openRun; empty runs are dropped) and
// heaps them by their front records. An error closes what was opened.
func OpenMerger(paths []string) (*Merger, error) {
	m := &Merger{}
	for _, p := range paths {
		rr, ok, err := openRun(p)
		if err != nil {
			m.Close()
			return nil, err
		}
		if ok {
			m.h = append(m.h, rr)
		}
	}
	for i := len(m.h)/2 - 1; i >= 0; i-- {
		siftDown(m.h, i)
	}
	return m, nil
}
