//ff:func feature=pipeline type=helper control=sequence
//ff:what packWriter.checkpointDue — whether to checkpoint now: every documents since the last, or the guard entered caution since it was last asked
package pipeline

// checkpointDue answers each entry into caution once.
func (w *packWriter) checkpointDue() bool {
	n := w.guard.Cautions()
	due := w.since >= w.every || n != w.cautions
	w.cautions = n
	return due
}
