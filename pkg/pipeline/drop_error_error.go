//ff:func feature=pipeline type=helper control=sequence
//ff:what DropError.Error — the entity and why its dump line does not decode
package pipeline

// Error names the entity and the decoding error.
func (e DropError) Error() string {
	return "dump line of " + e.ID + " does not decode: " + e.Err.Error()
}
