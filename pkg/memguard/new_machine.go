//ff:func feature=memguard type=helper control=sequence
//ff:what newMachine — a machine at Normal with every worker active
package memguard

// newMachine starts with workers workers active (at least 1) under l, which
// has its defaults.
func newMachine(l Limits, workers int) machine {
	workers = max(1, workers)
	return machine{l: l, workers: workers, active: workers, hold: l.Hold}
}
