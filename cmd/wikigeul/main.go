//ff:func feature=cli type=command control=sequence
//ff:what main — wikigeul entrypoint: run and print the error; exit 1 on failure, 3 when a strict run stopped at an anomaly, 4 when the memory guard stopped it
package main

import (
	"fmt"
	"os"
)

// main runs wikigeul (README) and exits with exitCode; a stop for memory also
// prints the command that goes on (exitCode).
func main() {
	err := run(os.Args[1:])
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, err)
	os.Exit(exitCode(err, os.Args, os.Stderr))
}
