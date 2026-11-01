//ff:func feature=dumpfile type=writer control=sequence
//ff:what saveState — save the sort state atomically: temporary file, fsync, rename, fsync the directory
package dumpfile

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// saveState writes st to dir/state.json so that a crash leaves the old or
// the new state, never a part: through state.json.tmp, synced, renamed, and
// the directory synced.
func saveState(dir string, st State) error {
	b, _ := json.Marshal(st) // plain fields: never fails
	tmp := filepath.Join(dir, "state.json.tmp")
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	if _, err := f.Write(b); err != nil {
		f.Close()
		return err
	}
	if err := f.Sync(); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	if err := os.Rename(tmp, filepath.Join(dir, "state.json")); err != nil {
		return err
	}
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	return d.Sync()
}
