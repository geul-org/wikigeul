//ff:func feature=memguard type=helper control=sequence
//ff:what fixtureSample — a 16 GiB machine with pct percent available and a 1 GiB process
package memguard

// fixtureSample makes floors easy to read: 30 caution, 15 danger, 35 recovered.
func fixtureSample(pct int64) Sample {
	const total = 16 << 30
	return Sample{RSS: 1 << 30, Total: total, Avail: total / 100 * pct}
}
