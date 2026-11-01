//ff:func feature=cli type=helper control=sequence
//ff:what dropError.Error — the entity and why its dump line does not decode
package main

// Error names the entity and the decoding error.
func (e dropError) Error() string {
	return "dump line of " + e.ID + " does not decode: " + e.Err.Error()
}
