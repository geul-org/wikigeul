//ff:func feature=pipeline type=helper control=iteration dimension=1
//ff:what FirstQ — the first n Q entities of an ID list; nil for no list
package pipeline

// FirstQ returns the first n Q jobs of ids, the entities dictSamples takes
// from a list (a list with none gives an empty, non-nil slice); no list
// (nil) gives nil.
func FirstQ(ids []PackJob, n int) []PackJob {
	if ids == nil {
		return nil
	}
	out := []PackJob{}
	for _, j := range ids {
		if j.Kind == 'Q' && len(out) < n {
			out = append(out, j)
		}
	}
	return out
}
