//ff:type feature=memguard type=model
//ff:what Level — how close the process stands to its memory limits: normal, caution (slow down), danger (stop at a safe point)
package memguard

// Level is the guard's state. Normal runs at full speed. Caution slows the
// work down and gives memory back, and is left again when memory recovers.
// Danger stops the run at its next safe point and is not left.
type Level int32

const (
	Normal Level = iota
	Caution
	Danger
)
