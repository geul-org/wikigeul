//ff:func feature=memguard type=helper control=iteration dimension=1
//ff:what ResumeCommand — the command line to continue a stopped run: the same arguments, shell-quoted, with -resume added when missing
package memguard

import (
	"regexp"
	"strings"
)

// shellPlain matches arguments the shell takes as they are.
var shellPlain = regexp.MustCompile(`^[A-Za-z0-9_@%+=:,./-]+$`)

// ResumeCommand prints args (os.Args) for a shell.
func ResumeCommand(args []string) string {
	var out []string
	resume := false
	for _, a := range args {
		resume = resume || a == "-resume" || a == "--resume"
		if !shellPlain.MatchString(a) {
			a = "'" + strings.ReplaceAll(a, "'", `'\''`) + "'"
		}
		out = append(out, a)
	}
	if !resume {
		out = append(out, "-resume")
	}
	return strings.Join(out, " ")
}
