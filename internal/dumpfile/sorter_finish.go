//ff:func feature=dumpfile type=writer control=sequence
//ff:what sorter.finish — at the end of the dump: last run, the file's SHA-256 checked against the expected one, state marked scanned
package dumpfile

import "fmt"

// finish writes the last run (spill), finishes the input (the whole file
// hashed) and compares its SHA-256 with the one the documents will record
// (Info.SHA256). A mismatch is an error and leaves the sort unscanned: the
// file is not the dump the provenance names.
func (s *sorter) finish() error {
	if err := s.spill(); err != nil {
		return err
	}
	sum, err := s.in.finish()
	if err != nil {
		return err
	}
	if sum != s.st.SHA256 {
		return fmt.Errorf("wikigeul: SHA-256 of %s is %s, not %s", s.c.Info.Name, sum, s.st.SHA256)
	}
	s.st.Scanned = true
	return saveState(s.c.Dir, *s.st)
}
