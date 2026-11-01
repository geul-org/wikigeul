//ff:func feature=dumpfile type=loader control=sequence
//ff:what runReader.close — release the run's decoder and file
package dumpfile

// close releases the zstd decoder and closes the run file.
func (rr *runReader) close() {
	rr.z.Close()
	rr.f.Close()
}
