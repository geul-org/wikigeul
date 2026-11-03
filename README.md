# wikigeul

`wikigeul` converts Wikidata entities into [GEUL](https://geul.org/) documents. It reads the Wikidata JSON dump file directly and publishes one `.geul` document per item or property — labels, descriptions, aliases, claims with qualifiers and references, and sitelinks in every language — as a pack set under `.well-known/geul/`.

It uses no database. To convert a dump loaded into PostgreSQL use `wikipg export-geul` from [wikipgeul](https://github.com/geul-org/wikipgeul), which runs the conversion code of this repository (`pkg/port`, `pkg/pipeline`) over the database: same document bytes for the same dump, commits, registry and `-recorded-at`.

All documents are published under one Issuer, `wikidata.geul.org` (`0x00000001`). The top byte of a document ID is the entity kind letter (`Q` = `0x51`, `P` = `0x50`) and the low 56 bits are its number, so Q42 is document `0x510000000000002a`.

## Install

```sh
go install github.com/geul-org/wikigeul/cmd/wikigeul@latest
```

## Requirements

- Go 1.25 or later.
- The dump file: `.json`, `.json.gz`, or `.json.bz2` with `lbzip2` (preferred) or `bzip2` on `PATH`, and its SHA-256 as `<file>.sha256` or `-dump-sha256`.
- The full git commit IDs of this tool and of the [gogeul](https://github.com/geul-org/gogeul) library it was built with; they are recorded in every document.
- An archive holding the LANGUAGE-SHORT-CODE codec document `0x00000000/codecs/0x000007a0` (the GEUL Language Registry), by default `-root` itself.

## Usage

```sh
wikigeul -dump wikidata-YYYYMMDD-all.json.bz2 -root ./archive -from 1 -to 200000000 \
  -commit "$(git rev-parse HEAD)" -lib-commit <gogeul commit> \
  -sort-dir ./wikigeul-dumpsort -report report.tsv
```

The documents go straight into the Issuer's pack set `<root>/.well-known/geul/0x00000001/` (zstd, with a dictionary trained on the run's first Q documents): every property of the dump, then the dump's items Q<from>…Q<to> (or the entities of `-ids`), converted by `-workers` goroutines and written in document key order as one stream (gogeul `pack.Stream`). Each document is version 1 with its version list and is also served by content hash under `0x00000001/hashes/` without a copy (addresses `0x00000001/docs/<document ID>/0x00000001.geul` and `…/versions.geul`). No document list (`0x00000001/docs/`, the name index) is written: for all of Wikidata it would be a single 4 GB document, and a document's name follows from its ID. The run checkpoints every `-checkpoint` documents; `-resume` continues an interrupted run, which must be the same run (commits, registry version, dump, range or ID list, `-recorded-at`). The set must not exist yet: a run writes a whole new set.

Each entity line is decoded into the `source.Entity` that `wikipg export-geul` reads after wikipgeul loaded that line (`source.DecodeDumpEntity`: the same orders, value forms including PostgreSQL's jsonb text, qualifiers-order, each reference occurrence's own snaks-order, sitelinks). The provenance comes from the file as wikipgeul would record it: `-dump-url` (default: the dumps.wikimedia.org URL from the file name's date), `-dump-sha256` (default: `<file>.sha256`; the file is checked against it), `-dump-time` (default: the file's modification time).

A dump is in no ID order, but the pack stream takes documents in key order. So the run first sorts the entity lines it converts (every property and Q<from>…Q<to>, or `-ids`) into zstd runs of `-sort-mem` MiB under `-sort-dir`, merged down to at most 64, then merges them into key order while converting. The sort checkpoints after every run and resumes by itself when the command is run again (a plain file at the exact byte offset; `.gz`/`.bz2` are decompressed again from the start and skipped up to it); the conversion checkpoints and resumes (`-resume`). Details, the table of what maps to what and the few known differences: [docs/dump.md](docs/dump.md).

| Flag | Meaning | Default |
|---|---|---|
| `-dump` | Wikidata JSON dump file (`.json`, `.json.gz`, `.json.bz2`) (required) | |
| `-root` | Archive root; documents go under `<root>/.well-known/geul/` | `.` |
| `-from`, `-to` | QID range | `1`, `1000` |
| `-commit` | Full commit ID of wikigeul, recorded in TRANSFORMED-BY (required) | |
| `-lib-commit` | Full commit ID of the gogeul library, recorded in TRANSFORMED-BY (required) | |
| `-report` | Write a per-entity TSV report to this path | none |
| `-workers` | Conversion goroutines; also the sort's decompressor and encoder threads | 70% of the CPUs |
| `-ids` | File of entity IDs (`Q…`/`P…`, one per line) instead of every property and the Q range | none |
| `-checkpoint` | Documents between checkpoints | `100000` |
| `-resume` | Continue the interrupted run from its last checkpoint | off |
| `-recorded-at` | RECORDED-AT of every version list (Unix seconds) | the run's start |
| `-samples` | Q documents the zstd dictionary is trained on | `2048` |
| `-strict` | Stop at the first anomaly (see below) after a checkpoint, exit code 3 | off |
| `-anomaly` | `-strict`: anomaly report path (TSV) | `wikigeul-anomaly.tsv` |
| `-allow` | `-strict`: also treat `kind:reason` as normal (`skipped:…`, `raw:…`, `warn:<rule>`); repeatable, not part of the run's identity | none |
| `-dump-url`, `-dump-sha256`, `-dump-time` | The dump's URL, file SHA-256 (hex) and download time (Unix seconds) recorded in every document | from the file name, `<file>.sha256`, the file's modification time |
| `-sort-dir` | Directory of the sort into key order (runs, `state.json`) | `wikigeul-dumpsort` |
| `-sort-mem` | MiB of entity JSON per sorted run | `1024` |
| `-mem-min-free`, `-mem-danger-free` | Percent of memory to keep free: below the first the run slows down and gives memory back, below the second (default half) it checkpoints and stops, exit code 4 | `30`, half |
| `-mem-max` | Cap on this process's RSS, e.g. `6GiB` | none |
| `-mem-interval`, `-mem-off` | Time between memory samples; turn the guard off | `1s`, on |
| `-langreg-root` | Archive to read the GEUL Language Registry (codec document `0x00000000/codecs/0x000007a0`) from | `-root` |
| `-langreg-version` | Version of that spec document to use | `0`: the newest when the run starts |

Memory guard: the run samples its own RSS and the memory available on the machine (and in its cgroup) throughout the sort, the conversion and the closing of the set. Below `-mem-min-free` it runs fewer workers and gives memory back, and goes up again when memory recovers. Below `-mem-danger-free`, or near `-mem-max`, it checkpoints at the next safe point and exits 4, printing `resume: <the same command> -resume`; the output is the same bytes as a run never stopped. Details: [docs/memory.md](docs/memory.md).

Strict mode (`-strict`): a run over all of Wikidata must not silently drop what it has never seen. Before a result is written, the writer looks for anomalies: a failed conversion or a panic, an entity of the ID list the dump does not hold, a skipped value other than a Lexeme (`L…`, forms, senses) or EntitySchema (`E…`) value, a value kept as original text for a reason not met in the validated runs (coordinate precision; time with timezone or bounds; time precision 0–5, millennium, century, Julian decade; BCE or 5+ digit decade; Y-prefixed year; day-precision date that is no day), a reference mismatch, and any `schema.Check` Fact on the document or its version list other than Warn E6. At the first anomalous entity it checkpoints everything before it, writes the anomaly report (`id`, `kind`, `property`, `value` (original, quoted), `reason`, `class`: 문법 부족 / 변환기 미지원 / 데이터 이상 / 기타) and exits 3. `-resume` starts again at that entity: add `-allow kind:reason` if the case is accepted as normal; if the porter must change, its commit changes and the run starts over. A dump line that does not decode as its entity (wikipgeul would have loaded none) is left out with a note on stderr, and in strict mode is the anomaly `decode` (`-allow decode:덤프 줄 해석 실패`). A panic is recovered into an error in every mode. After every checkpoint the run prints `progress <time> last=… docs=… session_docs=… rate=… skipped={…} raw={…}`.

Language codes: each Wikimedia language code is mapped to a BCP 47 tag; when the GEUL Language Registry numbers the tag (compared without case) the text's language is written as a 1-WORD LANGUAGE-SHORT-CODE, otherwise as a LANGUAGE-BCP47-CODE in canonical case. The registry is not built into the code: it is read once per run from the spec document's ENUMs through the gogeul resolve path, and the version used (`0x00000000/codecs/0x000007a0@n` and its content hash) is printed and recorded in the report's first line.

The report starts with a comment line recording the run (`# wikigeul -from … -to … -dump <file name> sha256:<hex> -commit … -lib-commit … -langreg 0x00000000/codecs/0x000007a0@n 0x<hash>`; `wikipg export-geul` writes the same line with `-schema … -dump-id …`), then a header and one line per ID in the range. `status` is `ok` (document written) or `missing` (no such entity in the dump). The other columns count what was written and what was skipped: `labels`, `descriptions`, `aliases`, `triples`, `qualifiers`, `sitelinks`, `badges`, `skipped`, `raw` (values kept as original text), `markers` (unknown/no-value markers), and the reference columns `sources`, `origin_qid`, `origin_absent`, `origin_local`, `origin_edges`, `source_qualifiers`, `ref_mismatch` (must be 0).

Edge IDs start from 1, as a run writes a new set. Lexemes (`L…`) are not converted.

## Documentation

Detailed rules (in Korean):

- [docs/conversion.md](docs/conversion.md) — document IDs, record order, value conversion, language codes, property datatypes, sitelinks
- [docs/provenance.md](docs/provenance.md) — document-level sources, TRANSFORMED-BY, conversion of Wikidata references
- [docs/dump.md](docs/dump.md) — reading the dump file: database ↔ dump JSON mapping, known differences, order (sort and merge), resume, flags, verification

The GEUL grammar itself is specified at [geul.org/grammar](https://geul.org/grammar/).

## Layout

```
cmd/wikigeul/      command: flags, the dump's provenance, the sort, jobs from the sorted dump
internal/dumpfile/ reading a dump file: decompression, entity lines, sorted runs and merge, checkpoints
pkg/source/        the entity model, one dump line as an entity (shared with wikipg export-geul)
pkg/port/          value conversion, document building, provenance (shared)
pkg/pipeline/      a conversion run whichever its input: flags, language registry, pack stream, workers, writer, checkpoints, strict mode, report (shared)
testdata/          the language code list of the injectivity test
```

## Development

```sh
go build ./... && go vet ./... && go test ./...
```

## License

Our own work in this repository is dedicated to the public domain under [CC0 1.0](LICENSE). Third-party components keep their own licenses; see [NOTICE](NOTICE). Wikidata content and the documents produced from it are CC0.
