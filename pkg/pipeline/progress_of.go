//ff:func feature=pipeline type=helper control=sequence
//ff:what progressOf — the identity a pack-mode run checkpoints (commits, registry, dump or dump file, range or ID list, RECORDED-AT), nothing done yet
package pipeline

// progressOf is the PackProgress of run c with ID list hash ids, at the start.
func progressOf(c Config, ids string) PackProgress {
	return PackProgress{Commit: c.Commit, LibCommit: c.LibCommit, LangDoc: c.LangDoc, Schema: c.Schema, Dump: c.Dump,
		From: c.From, To: c.To, IDs: ids, Source: c.Source, RecordedAt: c.RecordedAt}
}
