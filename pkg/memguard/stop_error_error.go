//ff:func feature=memguard type=helper control=sequence
//ff:what StopError.Error — the message of a stop: phase, reason, what to do
package memguard

import "fmt"

// Error says where and why the run stopped.
func (e StopError) Error() string {
	return fmt.Sprintf("memory guard: stopped in %s: %s; progress is saved, free memory and run the same command with -resume", e.Phase, e.Reason)
}
