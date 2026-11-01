//ff:func feature=cli type=helper control=selection
//ff:what keepEntity — which dump entities a run converts: the listed ones, or every property and the Q range (as the database path lists them)
package main

// keepEntity decides by letter and number which entities of the dump the run
// converts, as produceJobs lists them from the database: with an ID list,
// exactly the listed ones; else every property (up to 2^56−1, pageKind's
// bound) and the items Q<From>…Q<To>. Other kinds (Lexemes, Entity Schemas)
// are never converted.
func keepEntity(c config, ids []packJob) func(byte, uint64) bool {
	switch {
	case ids != nil:
		listed := map[packJob]bool{}
		for _, j := range ids {
			listed[packJob{Kind: j.Kind, Num: j.Num}] = true
		}
		return func(kind byte, num uint64) bool { return listed[packJob{Kind: kind, Num: num}] }
	default:
		return func(kind byte, num uint64) bool {
			return kind == 'P' && num <= ^uint64(0)>>8 || kind == 'Q' && num >= c.From && num <= c.To
		}
	}
}
