//ff:func feature=cli type=command control=sequence
//ff:what decodeJob — the entity of a dump job decoded from its line; a DropError when the line does not decode or names another entity
package main

import (
	"fmt"
	"strconv"

	"github.com/geul-org/wikigeul/pkg/pipeline"
	"github.com/geul-org/wikigeul/pkg/source"
)

// decodeJob decodes job's dump line (source.DecodeDumpEntity). A line
// wikipgeul could not decode either, or whose decoded ID is not the one the
// sort read from it (the line keys the sort by its first "id" member;
// encoding/json takes the last one and matches keys without case), is a
// DropError: wikipgeul would have loaded no such entity.
func decodeJob(job pipeline.PackJob) (source.Entity, bool, error) {
	id := string(job.Kind) + strconv.FormatUint(job.Num, 10)
	e, err := source.DecodeDumpEntity([]byte(job.Line))
	if err != nil {
		return source.Entity{}, false, pipeline.DropError{ID: id, Err: err}
	}
	if e.ID != id {
		return source.Entity{}, false, pipeline.DropError{ID: id, Err: fmt.Errorf("the line decodes as %s", e.ID)}
	}
	return e, true, nil
}
