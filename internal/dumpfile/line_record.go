//ff:func feature=dumpfile type=loader control=sequence
//ff:what lineRecord — one dump line as a Record when it holds an entity the filter keeps
package dumpfile

import "fmt"

// lineRecord takes dump line number line: its entity JSON (trimLine), the
// ID's letter and number (lineID, SplitID) and whether keep keeps it. A line
// without an entity is lineNone, one whose ID cannot be read or split is
// lineBad with the reason — wikipgeul loads no entity for it either.
func lineRecord(raw []byte, line int64, keep func(byte, uint64) bool) (Record, lineStatus, error) {
	b, ok := trimLine(raw)
	if !ok {
		return Record{}, lineNone, nil
	}
	id, err := lineID(b)
	if err != nil {
		return Record{}, lineBad, err
	}
	kind, num, ok := SplitID(id)
	if !ok {
		return Record{}, lineBad, fmt.Errorf("entity ID %q does not split into a letter and a number", id)
	}
	if !keep(kind, num) {
		return Record{}, lineSkipped, nil
	}
	return Record{Kind: kind, Num: num, Line: line, Data: b}, lineKept, nil
}
