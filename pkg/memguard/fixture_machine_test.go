//ff:func feature=memguard type=helper control=sequence
//ff:what fixtureMachine — a machine at the default limits
package memguard

// fixtureMachine has a 1 s interval, Lead 10 s, Hold 30 s, Cooldown 5 s.
func fixtureMachine(workers int) machine {
	return newMachine(Limits{}.withDefaults(), workers)
}
