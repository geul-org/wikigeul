//ff:func feature=dumpfile type=loader control=sequence
//ff:what input.finish — after the text ended: wait for the decompressor, hash what is left of the file, the file's SHA-256 in hex
package dumpfile

import (
	"encoding/hex"
	"io"
)

// finish is called when the text has ended: it waits for a decompressor
// program (which has then read its whole input), reads any bytes of a
// compressed file not yet read through the hash, and returns the SHA-256 of
// the whole file in hex.
func (in *input) finish() (string, error) {
	if in.cmd != nil {
		err := in.cmd.Wait()
		in.cmd = nil
		if err != nil {
			return "", err
		}
	}
	if !in.plain {
		if _, err := io.Copy(io.Discard, in.raw); err != nil {
			return "", err
		}
	}
	return hex.EncodeToString(in.h.Sum(nil)), nil
}
