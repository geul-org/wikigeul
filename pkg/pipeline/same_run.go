//ff:func feature=pipeline type=helper control=sequence
//ff:what sameRun — whether a checkpoint was made by the same run (commits, registry, dump, range or ID list, RECORDED-AT)
package pipeline

// sameRun compares the identity fields of a resumed checkpoint p with the run
// now (want); progress fields (Kind, Num, Docs, Report, Drained) are not compared.
func sameRun(p, want PackProgress) bool {
	want.Kind, want.Num, want.Docs, want.Report, want.Drained = p.Kind, p.Num, p.Docs, p.Report, p.Drained
	return p == want
}
