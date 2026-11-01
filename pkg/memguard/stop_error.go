//ff:type feature=memguard type=model
//ff:what StopError — a run stopped itself to keep memory: the phase it was in and why; the exit code is ExitCode
package memguard

// StopError is returned by the safe point that stopped the run, after the
// run's progress was saved. Nothing is lost: the same command with -resume
// goes on.
type StopError struct {
	Phase  string
	Reason string
}
