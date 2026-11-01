//ff:func feature=dumpfile type=writer control=sequence
//ff:what runWriter.close — finish the zstd stream, sync and close the run file
package dumpfile

import "errors"

// close ends the zstd stream and makes the file durable (fsync) before it
// is renamed and named in the sort state.
func (w *runWriter) close() error {
	err := w.z.Close()
	if err == nil {
		err = w.f.Sync()
	}
	return errors.Join(err, w.f.Close())
}
