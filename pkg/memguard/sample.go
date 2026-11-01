//ff:type feature=memguard type=model
//ff:what Sample — one reading: this process's RSS, the system's total and available memory, the tightest cgroup's limit and use
package memguard

// Sample is one reading, in bytes. Total and Avail are MemTotal and
// MemAvailable. CgroupLimit is the tightest memory limit over this process's
// cgroups (0: none) and CgroupUsed that cgroup's memory less its inactive
// file cache, which the kernel takes back first.
type Sample struct {
	RSS         int64
	Total       int64
	Avail       int64
	CgroupLimit int64
	CgroupUsed  int64
}
