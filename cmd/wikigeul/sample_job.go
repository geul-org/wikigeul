//ff:func feature=cli type=helper control=selection
//ff:what sampleJob — a dump job as a dictionary sample: one whose line does not decode is no entity, as the database path sees it
package main

// sampleJob makes j a dictionary sample the way the database path samples
// (dictSamples): a dump line that does not decode (decodeJob) is an entity
// wikipg would not have loaded. Taken from the list (listed) it stays, with
// no line, so buildAll skips it as missing; taken from the sorted dump it is
// not counted (false) and the next entity is sampled instead, as
// source.EntityNumbers would not list it.
func sampleJob(j packJob, listed bool) (packJob, bool) {
	_, _, err := decodeJob(j)
	switch {
	case j.Line == "" || err == nil:
		return j, true
	case listed:
		j.Line = ""
		return j, true
	default:
		return j, false
	}
}
