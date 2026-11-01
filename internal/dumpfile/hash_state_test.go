//ff:func feature=dumpfile type=helper control=sequence
//ff:what hashState — the saved SHA-256 state after some bytes, as sorter.spill saves it for a plain file
package dumpfile

import (
	"crypto/sha256"
	"encoding"
)

// hashState is the marshaled SHA-256 state after hashing s.
func hashState(s string) []byte {
	h := sha256.New()
	h.Write([]byte(s))
	b, _ := h.(encoding.BinaryMarshaler).MarshalBinary()
	return b
}
