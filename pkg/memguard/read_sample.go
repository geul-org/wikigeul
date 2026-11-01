//ff:func feature=memguard type=loader control=sequence
//ff:what readSample — one Sample from meminfo, statm and the cgroup; a cgroup that cannot be read is left out
package memguard

import "os"

// readSample reads the files. The cgroup is optional: when it cannot be read
// the sample has no cgroup limit.
func readSample(meminfo, statm string, cg cgroup) (Sample, error) {
	b, err := os.ReadFile(meminfo)
	if err != nil {
		return Sample{}, err
	}
	var s Sample
	if s.Total, s.Avail, err = parseMeminfo(b); err != nil {
		return Sample{}, err
	}
	if s.RSS, err = readRSS(statm); err != nil {
		return Sample{}, err
	}
	if used, err := cg.used(); err == nil && cg.Limit > 0 {
		s.CgroupLimit, s.CgroupUsed = cg.Limit, used
	}
	return s, nil
}
