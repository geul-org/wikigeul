//ff:type feature=dumpfile type=loader
//ff:what input — the dump being read: the file, its decompressed text, the SHA-256 of its bytes and the decompressor
package dumpfile

import (
	"hash"
	"io"
	"os"
	"os/exec"
)

// input is the dump file open for a sort. r is the decompressed JSON text.
// h hashes the file's bytes: for a plain file the lines as they are taken
// (sorter.line), so its state at any checkpoint matches State.Offset; for a
// compressed file every byte read from f, through raw. A .bz2 runs through
// an external program (cmd, its output out); a .gz through gz.
type input struct {
	f     *os.File
	raw   io.Reader
	r     io.Reader
	h     hash.Hash
	plain bool
	cmd   *exec.Cmd
	out   io.ReadCloser
	gz    io.Closer
}
