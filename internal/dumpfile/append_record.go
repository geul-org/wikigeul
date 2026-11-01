//ff:func feature=dumpfile type=writer control=sequence
//ff:what appendRecord — one record in run file format: header (kind, number, line, length) and the entity JSON
package dumpfile

import "encoding/binary"

// appendRecord appends r to buf as a run file stores it: kind, number, dump
// line and data length big-endian (recordHeaderSize bytes), then the data.
func appendRecord(buf []byte, r Record) []byte {
	buf = append(buf, r.Kind)
	buf = binary.BigEndian.AppendUint64(buf, r.Num)
	buf = binary.BigEndian.AppendUint64(buf, uint64(r.Line))
	buf = binary.BigEndian.AppendUint32(buf, uint32(len(r.Data)))
	return append(buf, r.Data...)
}
