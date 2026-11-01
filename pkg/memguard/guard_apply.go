//ff:func feature=memguard type=command control=sequence
//ff:what Guard.apply — carry out a decision: publish the numbers, stop first, set the workers, then give memory back and set the heap limit
package memguard

// apply acts on d. The stop is published before the workers change, so a
// gated worker wakes to it; memory is given back after the workers fell,
// so what they let go (encoders) goes with it.
func (g *Guard) apply(d decision) {
	if d.Level == Caution && g.level.Load() != int32(Caution) {
		g.cautions.Add(1)
	}
	g.level.Store(int32(d.Level))
	g.near.Store(d.Near)
	g.spare.Store(d.Press.Caution)
	g.budget.Store(min(d.Budget, g.base))
	if d.Reason != "" {
		g.reason = d.Reason
	}
	if d.Level == Danger && g.stop.CompareAndSwap(nil, &StopError{Phase: g.Phase(), Reason: g.reason}) {
		close(g.stopped)
	}
	g.setActive(d.Active)
	if d.Reclaim {
		g.reclaim()
	}
	g.setLimit(d.Budget)
}
