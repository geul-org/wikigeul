//ff:func feature=dumpfile type=writer control=sequence
//ff:what sorter.spill — write the held records as a sorted run and save the state: a sort checkpoint
package dumpfile

import (
	"encoding"
	"fmt"
	"path/filepath"
	"slices"
)

// spill sorts the held records (compareRecords) and writes them as the next
// run (writeRun), then saves the state with the run, the lines and bytes
// read and, for a plain file, the SHA-256 state at that offset — a
// checkpoint every kept line before which is in a run. It logs a progress
// line.
func (s *sorter) spill() error {
	slices.SortFunc(s.recs, compareRecords)
	if len(s.recs) > 0 {
		name := fmt.Sprintf("run-%06d.zst", s.st.Next)
		if err := writeRun(filepath.Join(s.c.Dir, name), s.recs, s.c.Threads); err != nil {
			return err
		}
		s.st.Runs = append(s.st.Runs, name)
		s.st.Next++
	}
	if s.in.plain {
		s.st.Hash, _ = s.in.h.(encoding.BinaryMarshaler).MarshalBinary() // sha256 state: never fails
	}
	s.recs, s.held, s.since = nil, 0, 0
	fmt.Fprintf(s.c.Log, "dump sort: line %d, %d bytes of text, kept %d, skipped %d, not read %d, %d runs\n",
		s.st.Lines, s.st.Offset, s.st.Kept, s.st.Skipped, s.st.Bad, len(s.st.Runs))
	return saveState(s.c.Dir, *s.st)
}
