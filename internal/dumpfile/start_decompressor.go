//ff:func feature=dumpfile type=loader control=selection
//ff:what startDecompressor — decompress a .gz in-process or a .bz2 through lbzip2/bzip2, reading the file through its hash
package dumpfile

import (
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/klauspost/compress/gzip"
)

// startDecompressor sets in.r to the text of the compressed file read
// through in.raw (so every byte is hashed). A .gz is read by a gzip reader
// (multistream, as pigz and gzip write). A .bz2 goes through the program
// bzip2Program on standard input — lbzip2 -dc -n threads decompresses even a
// single-stream file in parallel, as wikipg reads it — and in.cmd waits for
// it.
func startDecompressor(in *input, c SortConfig) error {
	switch {
	case strings.HasSuffix(c.Info.Name, ".gz"):
		zr, err := gzip.NewReader(in.raw)
		if err != nil {
			return err
		}
		in.r, in.gz = zr, zr
		return nil
	default:
		prog := bzip2Program()
		args := []string{"-dc"}
		if prog == "lbzip2" {
			args = append(args, "-n", strconv.Itoa(max(1, c.Threads)))
		}
		cmd := exec.Command(prog, args...)
		cmd.Stdin, cmd.Stderr = in.raw, os.Stderr
		out, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}
		in.r, in.out, in.cmd = out, out, cmd
		return cmd.Start()
	}
}
