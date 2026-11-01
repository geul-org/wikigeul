//ff:func feature=dumpfile type=command control=sequence
//ff:what Sort — sort the kept entity lines of a dump into at most mergeFanIn runs in document key order, resuming a saved sort
package dumpfile

import (
	"errors"
	"os"
)

// Sort reads the dump of c once, keeps the entity lines c.Keep accepts and
// writes them as sorted runs (scanDump), checks the file's SHA-256, then
// merges the runs down to at most mergeFanIn (reduceRuns), so one Merger
// over State.Runs yields the entities in document key order whatever order
// the dump has. Every step is saved in c.Dir (saveState): run again, it goes
// on from the last checkpoint. A directory holding the sort of another dump
// or filter is an error. A finished sort is returned as it is.
func Sort(c SortConfig) (State, error) {
	if err := os.MkdirAll(c.Dir, 0o755); err != nil {
		return State{}, err
	}
	st, ok, err := loadState(c.Dir)
	if err != nil {
		return State{}, err
	}
	want := State{Name: c.Info.Name, Size: c.Info.Size, ModTime: c.Info.ModTime, SHA256: c.Info.SHA256, Filter: c.Filter}
	if ok && !sameSort(st, want) {
		return State{}, errors.New("wikigeul: " + c.Dir + " holds the sort of another dump or entity range; remove it or give another -sort-dir")
	}
	if !ok {
		st = want
	}
	if !st.Scanned {
		if err := scanDump(c, &st); err != nil {
			return st, err
		}
	}
	return st, reduceRuns(c, &st)
}
