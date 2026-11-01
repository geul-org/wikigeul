//ff:type feature=memguard type=model
//ff:what Limits — what the guard keeps to: shares of memory left free, a cap on the process, and how often and how far ahead it looks
package memguard

import "time"

// Limits say how much memory a run may use. No byte count is built in: the
// floors are shares of the machine's memory (or of the cgroup's limit, the
// tighter of the two counts) and Max is whatever the caller declares.
//
// MinFree is the share kept free: below it the run slows down (Caution), and
// the Go heap is limited so that it stays above it. DangerFree is the share
// below which the run stops (Danger; 0 is half of MinFree). Gap is how far
// above MinFree memory must return before the guard goes up again (0 is
// defaultGap). Max caps the process's RSS in bytes (0: no cap): it is
// checked as a third limit with the floors capCaution, capDanger and
// capRecover, so the run stops before reaching it.
//
// Interval is the time between samples (0 is one second); it is a quarter of
// that while memory is below the recovery floor. Lead is how far ahead the
// trend is projected for Caution (0 is 10 intervals), Hold how long memory
// must stay recovered before each step up (0 is 30 intervals), Cooldown the
// time between two steps down (0 is 5 intervals), Confirm how many samples
// in a row Danger or a projected crossing needs (0 is defaultConfirm).
//
// Read replaces the system's numbers (tests); Meminfo is the file read for
// them otherwise (empty: /proc/meminfo). Off turns the guard off.
type Limits struct {
	Off        bool
	MinFree    float64
	DangerFree float64
	Gap        float64
	Max        int64
	Interval   time.Duration
	Lead       time.Duration
	Hold       time.Duration
	Cooldown   time.Duration
	Confirm    int
	Meminfo    string
	Read       Reader
}
