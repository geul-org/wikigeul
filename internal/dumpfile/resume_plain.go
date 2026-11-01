//ff:func feature=dumpfile type=loader control=sequence
//ff:what resumePlain — move a plain dump to the sort state's offset and restore the SHA-256 of the bytes before it
package dumpfile

import (
	"encoding"
	"io"
)

// resumePlain seeks the plain file of in to st.Offset and restores the
// SHA-256 state saved there (State.Hash). At offset 0 nothing is done.
func resumePlain(in *input, st State) error {
	if st.Offset == 0 {
		return nil
	}
	if err := in.h.(encoding.BinaryUnmarshaler).UnmarshalBinary(st.Hash); err != nil {
		return err
	}
	_, err := in.f.Seek(st.Offset, io.SeekStart)
	return err
}
