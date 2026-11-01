//ff:func feature=cli type=command control=sequence
//ff:what decodeJob — the entity of a dump job decoded from its line; a dropError when the line does not decode or names another entity
package main

import (
	"fmt"
	"strconv"

	"github.com/geul-org/wikigeul/internal/source"
)

// decodeJob decodes job's dump line (source.DecodeDumpEntity). A line
// wikipg could not decode either, or whose decoded ID is not the one the
// sort read from it (the line keys the sort by its first "id" member;
// encoding/json takes the last one and matches keys without case), is a
// dropError: wikipg would have loaded no such entity.
func decodeJob(job packJob) (source.Entity, bool, error) {
	id := string(job.Kind) + strconv.FormatUint(job.Num, 10)
	e, err := source.DecodeDumpEntity([]byte(job.Line))
	if err != nil {
		return source.Entity{}, false, dropError{ID: id, Err: err}
	}
	if e.ID != id {
		return source.Entity{}, false, dropError{ID: id, Err: fmt.Errorf("the line decodes as %s", e.ID)}
	}
	return e, true, nil
}
