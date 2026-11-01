//ff:func feature=pipeline type=command control=sequence
//ff:what packWriter.receive — the next result, false when the workers are done, or the checkpointed stop when the guard ends the run first
package pipeline

// receive waits on the results and on the guard.
func (w *packWriter) receive(results <-chan packResult) (packResult, bool, error) {
	select {
	case <-w.guard.Stopped():
		return packResult{}, false, w.memStop()
	case r, ok := <-results:
		return r, ok, nil
	}
}
