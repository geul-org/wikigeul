//ff:func feature=pipeline type=command control=sequence
//ff:what packWriter.drained — the workers are done: stop if the guard ended them, fail on a producer error, else close the stream
package pipeline

import "time"

// drained ends drain. Workers stopped by the guard close the results as well,
// and those results are not all of them.
func (w *packWriter) drained(perr *error) error {
	if w.guard.Err() != nil {
		return w.memStop()
	}
	if *perr != nil {
		return *perr
	}
	return w.finish(time.Now())
}
