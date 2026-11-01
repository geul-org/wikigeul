//ff:func feature=dumpfile type=writer control=iteration dimension=1
//ff:what writeRun — write sorted records as one run file, whole or not at all
package dumpfile

import "os"

// writeRun writes recs (already in run order) to path through a temporary
// file renamed at the end, so path is either absent or complete.
func writeRun(path string, recs []Record, threads int) error {
	w, err := newRunWriter(path+".tmp", threads)
	if err != nil {
		return err
	}
	for _, r := range recs {
		if err := w.put(r); err != nil {
			w.close()
			return err
		}
	}
	if err := w.close(); err != nil {
		return err
	}
	return os.Rename(path+".tmp", path)
}
