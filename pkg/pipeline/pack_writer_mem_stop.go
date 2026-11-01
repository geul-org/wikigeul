//ff:func feature=pipeline type=command control=sequence
//ff:what packWriter.memStop — stop for memory: checkpoint what was put, say where, return the guard's StopError (exit 4)
package pipeline

import "fmt"

// memStop saves everything put so far with its report and progress. Results
// not yet put are dropped: -resume converts them again, and the output does
// not depend on when that happens.
func (w *packWriter) memStop() error {
	if err := w.checkpoint(); err != nil {
		return fmt.Errorf("memory stop: checkpoint failed: %w", err)
	}
	fmt.Printf("pack: stopped for memory after %c%d, %d documents\n", w.prog.Kind, w.prog.Num, w.prog.Docs)
	return w.guard.Err()
}
