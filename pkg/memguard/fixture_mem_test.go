//ff:type feature=memguard type=model
//ff:what fixtureMem — what a test guard sees and does: the next sample or error, phases read, memory given back, limits set
package memguard

// fixtureMem stands in for the system.
type fixtureMem struct {
	s        Sample
	err      error
	phases   []string
	reclaims int
	limits   []int64
}
