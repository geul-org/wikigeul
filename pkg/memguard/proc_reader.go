//ff:func feature=memguard type=loader control=sequence
//ff:what procReader — the Reader of a Linux system: the cgroup found once, then meminfo, statm and the cgroup read for each sample
package memguard

// procReader reads meminfo (empty: /proc/meminfo) and this process's statm
// and cgroup.
func procReader(meminfo string) Reader {
	if meminfo == "" {
		meminfo = procMeminfo
	}
	cg := findCgroup(procCgroup, cgroupRoot)
	return func(string) (Sample, error) { return readSample(meminfo, procStatm, cg) }
}
