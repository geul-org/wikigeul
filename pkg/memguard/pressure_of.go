//ff:func feature=memguard type=helper control=sequence
//ff:what pressureOf — a sample against the limits: the system's memory, the cgroup's when it has a limit, the cap when set
package memguard

// pressureOf measures s against l: the free shares of the machine and of a
// limited cgroup, and the room left under Max.
func pressureOf(l Limits, s Sample) pressure {
	p := slackOf("system", s.Avail, s.Total, l.MinFree, l.DangerFree, l.MinFree+l.Gap)
	if s.CgroupLimit > 0 {
		p = p.tighter(slackOf("cgroup", s.CgroupLimit-s.CgroupUsed, s.CgroupLimit, l.MinFree, l.DangerFree, l.MinFree+l.Gap))
	}
	if l.Max > 0 {
		p = p.tighter(slackOf("max", l.Max-s.RSS, l.Max, capCaution, capDanger, capRecover))
	}
	return p
}
