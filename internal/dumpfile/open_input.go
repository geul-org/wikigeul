//ff:func feature=dumpfile type=loader control=selection
//ff:what openInput — open the dump at a sort state's offset: a plain file seeked with its hash state, a compressed one decompressed again from the start
package dumpfile

import (
	"crypto/sha256"
	"io"
	"os"
	"strings"
)

// openInput opens the dump of c and positions its text at st.Offset. A
// plain file (any name but .gz, .bz2) seeks there and restores the SHA-256
// state saved with it (resumePlain). A compressed file cannot be entered in
// the middle — a gzip member and a bzip2 stream have no byte-aligned restart
// points the decompressors expose — so it is decompressed from the start
// again (and hashed again) and the first st.Offset bytes of text are
// discarded.
func openInput(c SortConfig, st State) (*input, error) {
	f, err := os.Open(c.Info.Path)
	if err != nil {
		return nil, err
	}
	in := &input{f: f, h: sha256.New()}
	switch {
	case !strings.HasSuffix(c.Info.Name, ".gz") && !strings.HasSuffix(c.Info.Name, ".bz2"):
		in.plain, in.r = true, f
		err = resumePlain(in, st)
	default:
		in.raw = io.TeeReader(f, in.h)
		err = startDecompressor(in, c)
		if err == nil {
			_, err = io.CopyN(io.Discard, in.r, st.Offset)
		}
	}
	if err != nil {
		in.close()
		return nil, err
	}
	return in, nil
}
