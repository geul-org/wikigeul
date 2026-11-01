//ff:func feature=dumpfile type=command control=iteration dimension=1
//ff:what scanDump — read the dump from the sort state's offset to its end, writing a sorted run whenever enough is held
package dumpfile

import (
	"bufio"
	"errors"
	"io"
)

// scanDump reads the dump of c line by line from st.Offset (openInput),
// taking each line (sorter.line, lines of any length whole). When a run is
// due (sorter.due) it writes it and checkpoints (sorter.spill), and a run the
// guard stopped ends there; at the end it writes the last run and checks the
// file's SHA-256 (sorter.finish).
func scanDump(c SortConfig, st *State) error {
	c.Guard.SetPhase("sort:scan")
	in, err := openInput(c, *st)
	if err != nil {
		return err
	}
	defer in.close()
	s := &sorter{c: c, st: st, in: in}
	br := bufio.NewReaderSize(in.r, 1<<20)
	for {
		raw, err := br.ReadBytes('\n')
		if len(raw) > 0 {
			s.line(raw)
		}
		if err == nil && s.due() {
			err = s.spill()
		}
		if err == nil {
			err = c.Guard.Err()
		}
		if errors.Is(err, io.EOF) {
			return s.finish()
		}
		if err != nil {
			return err
		}
	}
}
