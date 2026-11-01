//ff:func feature=dumpfile type=writer control=iteration dimension=1
//ff:what mergeRuns — merge sorted runs into one run file, whole or not at all
package dumpfile

import "os"

// mergeRuns writes every record of the runs ins, in run order (Merger), to
// the run out through a temporary file renamed at the end.
func mergeRuns(out string, ins []string, threads int) error {
	m, err := OpenMerger(ins)
	if err != nil {
		return err
	}
	defer m.Close()
	w, err := newRunWriter(out+".tmp", threads)
	if err != nil {
		return err
	}
	for {
		r, ok, err := m.Next()
		if err == nil && ok {
			err = w.put(r)
		}
		if err != nil {
			w.close()
			return err
		}
		if !ok {
			break
		}
	}
	if err := w.close(); err != nil {
		return err
	}
	return os.Rename(out+".tmp", out)
}
