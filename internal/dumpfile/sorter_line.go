//ff:func feature=dumpfile type=loader control=selection
//ff:what sorter.line — take one dump line: count it, hash it (plain file), hold it when kept, log it when unreadable
package dumpfile

import "fmt"

// line counts raw as the next dump line and its bytes as read (a plain
// file's bytes are hashed here, so the hash stops exactly at the offset),
// then holds its record when the filter keeps it (lineRecord) and counts
// it as kept, skipped or unreadable; an unreadable line is logged with its
// number.
func (s *sorter) line(raw []byte) {
	s.st.Lines++
	s.st.Offset += int64(len(raw))
	s.since += int64(len(raw))
	if s.in.plain {
		s.in.h.Write(raw)
	}
	r, status, err := lineRecord(raw, s.st.Lines, s.c.Keep)
	switch status {
	case lineKept:
		s.recs = append(s.recs, r)
		s.held += int64(len(r.Data))
		s.st.Kept++
	case lineSkipped:
		s.st.Skipped++
	case lineBad:
		s.st.Bad++
		fmt.Fprintf(s.c.Log, "dump line %d not read: %v\n", s.st.Lines, err)
	}
}
