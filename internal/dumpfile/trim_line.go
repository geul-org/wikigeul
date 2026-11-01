//ff:func feature=dumpfile type=loader control=sequence
//ff:what trimLine — a dump line without its trailing comma and newline; false for the array brackets and blank lines
package dumpfile

import "bytes"

// trimLine returns the entity JSON of one dump line as wikipg takes it
// (wikipg dump.TrimLine): without the newline, carriage return and the comma
// between array elements, trimmed. The "[" and "]" lines and blank lines
// carry no entity and are false.
func trimLine(line []byte) ([]byte, bool) {
	b := bytes.TrimRight(line, "\r\n")
	b = bytes.TrimSuffix(b, []byte(","))
	b = bytes.TrimSpace(b)
	if len(b) == 0 || bytes.Equal(b, []byte("[")) || bytes.Equal(b, []byte("]")) {
		return nil, false
	}
	return b, true
}
