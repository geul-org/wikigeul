//ff:func feature=cli type=helper control=selection
//ff:what exitCode — the exit code of a failed run: 3 for a strict stop at an anomaly, 4 for a stop by the memory guard (telling the command that resumes), else 1
package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/geul-org/wikigeul/pkg/memguard"
	"github.com/geul-org/wikigeul/pkg/pipeline"
)

// exitCode says why the run ended. A stop by the guard writes the resume
// command to w, built from args (os.Args).
func exitCode(err error, args []string, w io.Writer) int {
	switch {
	case errors.As(err, &pipeline.AnomalyError{}):
		return 3
	case errors.As(err, &memguard.StopError{}):
		fmt.Fprintln(w, "resume:", memguard.ResumeCommand(args))
		return memguard.ExitCode
	default:
		return 1
	}
}
