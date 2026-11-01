//ff:func feature=dumpfile type=loader control=sequence
//ff:what loadState — the saved sort state of a directory; false when there is none
package dumpfile

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

// loadState reads dir/state.json (saveState). No file is false with no
// error.
func loadState(dir string) (State, bool, error) {
	b, err := os.ReadFile(filepath.Join(dir, "state.json"))
	if errors.Is(err, fs.ErrNotExist) {
		return State{}, false, nil
	}
	if err != nil {
		return State{}, false, err
	}
	var st State
	return st, true, json.Unmarshal(b, &st)
}
