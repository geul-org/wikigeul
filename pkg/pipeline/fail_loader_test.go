//ff:func feature=pipeline type=helper control=sequence
//ff:what failLoader — test fixture: a Loader whose every read fails
package pipeline

import (
	"errors"

	"github.com/geul-org/wikigeul/pkg/source"
)

// failLoader stands for an input that cannot be read.
func failLoader(PackJob) (source.Entity, bool, error) {
	return source.Entity{}, false, errors.New("cannot read")
}
