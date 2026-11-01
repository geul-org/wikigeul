//ff:func feature=cli type=helper control=sequence
//ff:what fixturePressedReader — a reader with plenty of memory, and almost none from the moment the run reaches phase
package main

import (
	"github.com/geul-org/wikigeul/pkg/memguard"
)

// fixturePressedReader is called by one guard, under its tick lock.
func fixturePressedReader(phase string) memguard.Reader {
	return func(now string) (memguard.Sample, error) {
		s := memguard.Sample{RSS: 1 << 30, Total: 1 << 40, Avail: 1 << 39}
		if now == phase {
			s.Avail = 1 << 30
		}
		return s, nil
	}
}
