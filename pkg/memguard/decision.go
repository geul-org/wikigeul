//ff:type feature=memguard type=model
//ff:what decision — what the state machine concludes from a sample: level, active workers, whether to give memory back, the process's budget, why
package memguard

// decision is one step's outcome. Active is how many workers may run (0 in
// danger), Reclaim asks for a collection and memory returned to the system.
// Budget is the RSS this process may reach with the caution floors kept,
// Press the sample against the limits, Fall the trend (bytes a second the
// slack is losing), Reason set when something changed.
type decision struct {
	Level   Level
	Active  int
	Reclaim bool
	Budget  int64
	Press   pressure
	Fall    float64
	Near    bool
	Reason  string
}
