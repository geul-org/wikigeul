//ff:func feature=dumpfile type=loader control=sequence
//ff:what input.close — stop a running decompressor and close the dump file
package dumpfile

// close stops a decompressor program still running (an interrupted sort):
// its output is closed, the process killed and waited for. Then the gzip
// reader and the file are closed.
func (in *input) close() {
	if in.cmd != nil {
		in.out.Close()
		in.cmd.Process.Kill()
		in.cmd.Wait()
		in.cmd = nil
	}
	if in.gz != nil {
		in.gz.Close()
	}
	in.f.Close()
}
