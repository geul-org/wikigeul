//ff:func feature=pipeline type=command control=sequence
//ff:what packWriter.finish — checkpoint that every entity is in, close the stream (manifest, latest.txt), printing what was written
package pipeline

import (
	"fmt"
	"time"
)

// finish checkpoints the report and the progress, marked Drained, and closes
// the stream, then prints the generation, documents, records and bytes. Close
// keeps its own progress but not the caller's, and the guard can stop it: the
// checkpoint gives a resumed run the report, the counts and the knowledge that
// no entity is left. Checkpoints do not change the bytes written.
func (w *packWriter) finish(start time.Time) error {
	w.prog.Drained = true
	if err := w.checkpoint(); err != nil {
		return err
	}
	res, err := w.s.Close()
	if err != nil {
		return err
	}
	fmt.Printf("pack: generation %d manifest %s, %d documents, %d records (%d zstd), %d objects, %d bytes\n",
		res.Generation, res.Manifest, w.prog.Docs, res.Records, res.Compressed, len(res.Objects), res.Bytes)
	return nil
}
