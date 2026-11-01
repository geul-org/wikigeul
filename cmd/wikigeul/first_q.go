//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what firstQ — the first n Q entities of an ID list; nil for no list
package main

// firstQ returns the first n Q jobs of ids, the entities dictSamples takes
// from a list (a list with none gives an empty, non-nil slice); no list
// (nil) gives nil.
func firstQ(ids []packJob, n int) []packJob {
	if ids == nil {
		return nil
	}
	out := []packJob{}
	for _, j := range ids {
		if j.Kind == 'Q' && len(out) < n {
			out = append(out, j)
		}
	}
	return out
}
