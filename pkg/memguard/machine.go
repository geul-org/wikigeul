//ff:type feature=memguard type=model
//ff:what machine — the guard's state: level, active workers, trend, confirmation counts and the times of the last steps; pure, no clock or system of its own
package memguard

import "time"

// machine turns samples into decisions (step); time and memory come in as
// arguments, so its transitions can be tested with neither. level and active
// are the state; last, prev, fall and rate the trend (rate unsmoothed); dangerN and cautionN count
// samples in a row with a crossing; calmSince is when memory last came above
// the recovery floor, lastShrink and lastRaise the last steps, hold the calm
// needed for a step up, longer after a raise that went wrong. ramp is set while
// the workers climb from one (Guard.Ramp); rssAt is the RSS at the last raise
// and cost the most RSS a raise has been seen to add.
type machine struct {
	l          Limits
	workers    int
	level      Level
	active     int
	last       time.Time
	prev       int64
	fall       float64
	rate       float64
	dangerN    int
	cautionN   int
	calmSince  time.Time
	lastShrink time.Time
	lastRaise  time.Time
	hold       time.Duration
	ramp       bool
	rssAt      int64
	cost       int64
}
