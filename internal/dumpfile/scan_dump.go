//ff:func feature=dumpfile type=command control=iteration dimension=1
//ff:what scanDump — read the dump from the sort state's offset to its end, writing a sorted run whenever enough is held
package dumpfile

import (
	"bufio"
	"errors"
	"io"
)

// scanDump reads the dump of c line by line from st.Offset (openInput),
// taking each line (sorter.line, lines of any length whole). When c.Mem
// bytes of entity JSON are held, or checkpointBytes of text were read, it
// writes a run and checkpoints (sorter.spill); at the end it writes the last
// run and checks the file's SHA-256 (sorter.finish).
func scanDump(c SortConfig, st *State) error {
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
		if err == nil && (s.held >= c.Mem || s.since >= checkpointBytes) {
			err = s.spill()
		}
		if errors.Is(err, io.EOF) {
			return s.finish()
		}
		if err != nil {
			return err
		}
	}
}
