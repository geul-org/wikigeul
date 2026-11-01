//ff:func feature=dumpfile type=writer control=sequence
//ff:what runWriter.put — write one record to the run
package dumpfile

// put writes r (appendRecord) to the run.
func (w *runWriter) put(r Record) error {
	w.buf = appendRecord(w.buf[:0], r)
	_, err := w.z.Write(w.buf)
	return err
}
