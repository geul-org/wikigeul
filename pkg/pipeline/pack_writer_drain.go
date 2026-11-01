//ff:func feature=pipeline type=command control=iteration dimension=1
//ff:what packWriter.drain — put every result in job order until the workers are done or the guard stops the run, then, if the producer did not fail, close the stream
package pipeline

// drain takes every result (take) until the workers are done; a failed put
// stops it. So does the guard: the writer waits on both (receive) and stops
// at once even with a long job running. A producer error (*perr, set before
// jobs closed, so before results closed) fails the run instead of closing an
// incomplete set; otherwise the stream is finished (drained).
func (w *packWriter) drain(results <-chan packResult, perr *error) error {
	for {
		r, ok, err := w.receive(results)
		if err != nil {
			return err
		}
		if !ok {
			return w.drained(perr)
		}
		if err := w.take(r); err != nil {
			return err
		}
	}
}
