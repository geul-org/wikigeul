package memguard

const (
	// ExitCode is the exit code of a run stopped by the guard (StopError). A
	// strict run stopped at an anomaly exits 3, a failed run 1.
	ExitCode = 4
	// DefaultMinFree is the share of memory left free by default.
	DefaultMinFree = 0.30
	// defaultGap is how far above the caution floor memory must come back
	// before the guard goes up again, as a share of memory.
	defaultGap = 0.05
	// defaultConfirm is how many samples in a row a projected crossing, or
	// danger, must last: one sample can be a burst.
	defaultConfirm = 2
	// capCaution, capDanger and capRecover are the shares of Limits.Max left
	// free at which caution begins, danger begins and caution ends.
	capCaution = 0.20
	capDanger  = 0.05
	capRecover = 0.30
	// maxHoldFactor bounds how far a raise that went wrong stretches Hold.
	maxHoldFactor = 8
	// logEverySamples is how many sampling intervals a quiet guard stays silent.
	logEverySamples = 60
	// shareFloor bounds Share from below: want/shareFloor.
	shareFloor = 16
	// mib is bytes per MiB, in log lines.
	mib = 1 << 20
	// Paths the guard reads by default.
	procMeminfo = "/proc/meminfo"
	procStatm   = "/proc/self/statm"
	procStatus  = "/proc/self/status"
	procCgroup  = "/proc/self/cgroup"
	cgroupRoot  = "/sys/fs/cgroup"
)
