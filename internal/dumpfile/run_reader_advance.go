//ff:func feature=dumpfile type=loader control=sequence
//ff:what runReader.advance — read the run's next record into cur; false at the end of the run
package dumpfile

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// advance reads the next record (appendRecord's format) into cur, with
// newly allocated data. The end of the run is false; a run that ends inside
// a record is corrupt.
func (rr *runReader) advance() (bool, error) {
	var h [recordHeaderSize]byte
	if _, err := io.ReadFull(rr.br, h[:]); err != nil {
		if errors.Is(err, io.EOF) {
			return false, nil
		}
		return false, fmt.Errorf("wikigeul: run %s: %w", rr.f.Name(), err)
	}
	data := make([]byte, binary.BigEndian.Uint32(h[17:]))
	if _, err := io.ReadFull(rr.br, data); err != nil {
		return false, fmt.Errorf("wikigeul: run %s: record cut short: %w", rr.f.Name(), err)
	}
	rr.cur = Record{Kind: h[0], Num: binary.BigEndian.Uint64(h[1:]), Line: int64(binary.BigEndian.Uint64(h[9:])), Data: data}
	return true, nil
}
