//ff:func feature=cli type=command control=sequence
//ff:what run — parse the command line and convert the dump file
package main

// run converts the dump file args name (runDump). The README lists the
// flags.
func run(args []string) error {
	c, err := parseFlags(args)
	if err != nil {
		return err
	}
	return runDump(c)
}
