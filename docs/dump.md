# 덤프 직접 입력

wikigeul은 위키데이터 JSON 덤프 파일을 바로 읽어 GEUL 문서를 만든다. PostgreSQL 적재본에서 만드는 일은 wikipgeul `export-geul`이 맡고, 변환 코드(`port.BuildDoc`, `pkg/pipeline`)는 이 저장소의 공개 패키지 하나를 함께 쓴다. 두 경로는 입력만 다르다. 둘 다 개체 하나를 `source.Entity`로 만들어 넘긴다.

- DB 경로(wikipgeul `export-geul`): `pgsource.LoadEntity`(SQL 질의).
- 덤프 경로(wikigeul): `source.DecodeDumpEntity`(덤프 한 줄 → wikipgeul가 그 줄을 적재한 뒤 DB 경로가 읽었을 `source.Entity`).
- 입력의 차이는 `pipeline.Loader` 하나로 모인다. 두 저장소가 각자의 Loader와 작업 생성을 주고, pack 쓰기(작업자, 순서 복원, 체크포인트, 엄격 모드, TRANSFORMED-BY 확인)는 `pipeline.WritePack` 한 벌이다.

목표는 같은 덤프·같은 커밋·같은 레지스트리 판·같은 `-recorded-at`이면 개체마다 DB 경로와 바이트 단위로 같은 문서다. 아래 "대응하지 않는 것"이 예외의 전부다.

```sh
wikigeul -dump /path/to/wikidata-YYYYMMDD-all.json.bz2 \
  -root ./archive -from 1 -to 200000000 \
  -commit "$(git rev-parse HEAD)" -lib-commit <gogeul commit> \
  -recorded-at <Unix 초> -sort-dir /path/to/sortdir -report report.tsv
```

## 입력

- `.json`(평문), `.json.gz`(gzip, 멀티 멤버 포함, 프로세스 안에서 klauspost/compress), `.json.bz2`(`lbzip2 -dc -n <workers>`, 없으면 `bzip2 -dc`, wikipgeul와 같다). 이름의 접미사로 고른다. 그 밖의 이름은 평문이다.
- 줄 처리는 wikipgeul와 같다. 1번 줄은 `[`이다. 줄 끝 `\r\n`과 쉼표를 떼고, `[`·`]`·빈 줄은 개체가 아니다.
- 범위는 DB 경로와 같다. 항목(Q)과 속성(P)만 옮기고, 그 밖의 종류(Lexeme L, EntitySchema E 등)는 건너뛴다. Lexeme 덤프(wikipgeul `-dump-id 2`)도 옮기지 않는다.
- 파일 전체의 SHA-256을 읽으면서 계산한다. 정렬이 끝날 때 문서에 적을 SHA-256과 다르면 변환을 시작하지 않는다.

## PG 경로 ↔ 덤프 JSON 대응표

DB 경로가 읽는 것(wikipgeul `pgsource.queriesFor`)은 wikipgeul가 덤프 줄에서 만든 행이다(wikipgeul `internal/load`). 덤프 경로는 같은 규칙을 줄에 바로 적용한다.

| 정보 | DB 경로 (wikipgeul 적재 → wikigeul 질의) | 덤프 경로 (`source.DecodeDumpEntity`) |
|---|---|---|
| 개체가 있는지 | `entity` 행이 있다(줄이 디코드되고 ID가 `SplitID`를 통과) | 줄의 `id`(`dumpfile.lineID`)가 `dumpfile.SplitID`를 통과하고 줄이 디코드된다 |
| 디코드 | `encoding/json` → wikipgeul `dump.Entity`(맵, 대소문자 무시 키 대응, 중복 키는 마지막) | 같은 필드·타입의 구조체(`dumpEntity` 등)로 `encoding/json`. Lexeme 필드만 뺐다(항목·속성 줄에는 없다) |
| datatype | `entity.datatype`(빈 값 NULL → `''`) | `datatype` |
| lastrevid | `entity.lastrevid`(0 → NULL → 0) | `lastrevid` |
| modified | `floor(extract(epoch FROM entity.modified))`, 없으면 0 | `modified`를 RFC 3339로 읽어 Unix 초(내림), 없으면 0 |
| 레이블·설명 | `label`·`description`, `language COLLATE "C"` 순. NUL이 있으면 `value_json`에서 원문 | `labels`·`descriptions` 맵 키(언어) 바이트 순, `value` 원문 |
| 별칭 | `alias`, 언어 바이트 순 → `ordinal` | `aliases` 키 바이트 순 → 배열 순 |
| claim 순서 | `statement_id` = 속성 ID 문자열 순, 같은 속성 안은 claims 배열 순 | `claims` 키를 문자열 정렬, 배열 순(`dumpClaims`). 번호는 개체 안 0부터(문서에 나오지 않는다) |
| claim 속성 | `statement.property` = mainsnak의 `property` | `mainsnak.property` |
| rank, snak 종류 | `rank::text`, `snaktype::text` | `rank`, `mainsnak.snaktype` |
| datatype(snak) | `coalesce(datatype, '')` | `mainsnak.datatype` |
| 값 종류 | `coalesce(value_type, '')` = `datavalue.type` | `datavalue.type`(datavalue 없으면 빈 값) |
| 값 | `coalesce(value_entity, value #>> '{}', '')`, `value_json`이 있으면 `exactValue` | `dumpSnakValue`: 항목 값에 `id`가 있으면 그 ID. `value` 원문에 `\u0000`이 있으면 wikipgeul처럼 원문 텍스트로 `exactValue`. 그 밖은 `jsonbText` |
| jsonb 출력 | PostgreSQL이 jsonb로 저장했다가 텍스트로 내놓은 것 | `jsonbText`가 재현: 문자열은 따옴표 없이, `null`은 빈 값, 객체 키는 길이 → 바이트 순(중복 키는 마지막), 구분자 `", "`·`": "`, 문자열 이스케이프는 `escape_json`(`\b\f\n\r\t\"\\`, 그 밖의 0x20 미만은 `\u00xx` 소문자, `/`·DEL·비ASCII는 그대로), 수는 numeric 출력(지수 없음, 소수 자릿수 = 원문 소수 자릿수 − 지수, 최소 0; `1.0e-5` → `0.000010`, `-0` → `0`) |
| qualifier 순서 | `qualifier.ordinal` = `qualifiers-order` 순, 거기 빠진 속성은 정렬해 뒤에, 속성 안은 배열 순 | 같은 규칙(`orderedSnaks`, wikipgeul와 같은 코드). order에 같은 속성이 두 번 있으면 두 번 놓이는 것까지 같다 |
| 출처 순서 | `statement_reference.ordinal` = `references` 배열 순 | `references` 배열 순 |
| 출처 snak | `reference_snak`(hash의 덤프상 첫 출현 목록) + `reference_snak_order`(출현별 순열) | 그 출현 자신의 `snaks`를 `snaks-order` 순(`orderedSnaks`) |
| snak 없는 출처 | `statement_reference ⋈ reference_snak`에 행이 없어 빠진다 | 빠진다(`dumpReferences`) |
| sitelink | `sitelink`, `site COLLATE "C"` 순, `title`(+`title_json`), `badges`(빈 배열 `{}`) | `sitelinks` 맵 키 바이트 순, `title` 원문, `badges`(없으면 빈 목록) |
| 덤프 출처(문서 HEAD) | `dump.url`, `dump.sha256`, `dump.downloaded_at` | `dumpfile.InfoOf`: `-dump-url`, 없으면 `https://dumps.wikimedia.org/wikidatawiki/entities/<파일 이름의 YYYYMMDD>/<파일 이름>`. `-dump-sha256`, 없으면 `<파일>.sha256`(sha256sum 형식). `-dump-time`, 없으면 파일 수정 시각. 기본값 규칙은 wikipgeul `dumpInfoOf`와 같다 |
| 출처 순서 방식 | `reference_snak_order`가 있으면 출현별 순서(`RefOrder`) | 항상 출현별 순서(`RefOrder = true`) |

## 대응하지 않는 것

DB 적재 과정에서만 생기거나 덤프 경로가 다르게 다루는 정보다.

1. **statement_id·줄 번호·`entity_raw`·`load_error`.** wikipgeul가 매기거나 남긴다. 문서에 나오지 않아 결과에 영향이 없다. 덤프 경로의 claim 번호는 개체 안 0부터다.
2. **적재 도구 확인.** DB 경로는 `dump.loader`가 wikipgeul가 아니면 거부한다. 덤프 경로에는 적재 도구가 없다. 대신 파일 SHA-256을 검증한다.
3. **`dump` 행의 출처 값.** wikipgeul 적재 때 `-url`을 줬거나, 파일을 옮겨 수정 시각이 바뀌었으면 기본값이 DB와 다르다. 같은 문서를 원하면 DB의 `dump` 행 값을 `-dump-url`·`-dump-sha256`·`-dump-time`(`floor(extract(epoch FROM downloaded_at))`)으로 준다. 압축 형식이 다른 파일(.gz와 .bz2)은 SHA-256이 달라 HEAD가 다르다. DB에 적재한 바로 그 파일을 쓴다. `dump_id`, `size`, `md5`, `sha1`, `dump_date`, `published_at`은 문서에 쓰지 않는다.
4. **출처 hash가 같은데 내용이 다른 출현.** DB는 hash마다 덤프상 첫 출현의 snak 목록 하나만 둔다(`reference_snak`). 순서만 다른 출현은 `reference_snak_order` 순열로 자기 순서를 되찾는다. 내용이 다른 출현은 순열이 없어 첫 출현의 내용으로 읽힌다. 덤프 경로는 늘 그 출현 자신의 내용이다. wikipgeul `build-reference-order` 보고의 `content` 수가 0이면 차이가 없다. 또 DB 경로가 `reference_snak_order` 없이 돌았으면(실행 첫 줄 `reference snak order: … false`) 순서만 다른 출현도 다르다.
5. **wikipgeul `-langs`.** 적재 때 언어를 골랐으면 DB에는 그 언어만 있다. 덤프 경로는 모든 언어를 읽는다(`-langs all`과 같다).
6. **PostgreSQL이 받지 못하는 값.** wikipgeul에서는 적재가 실패하거나 행이 빠진다. 덤프 경로는 Go의 디코드대로 옮긴다.
   - jsonb가 거부하는 짝 없는 서로게이트 `\ud800`(Go는 U+FFFD)
   - 잘못된 UTF-8
   - `value`가 없는 datavalue(빈 값)
   - enum 밖의 rank·snaktype
   - statement 2^20개 초과

   20261019 적재는 오류 0이었다.
7. **같은 ID가 두 줄.** DB는 PK `(entity_id, dump_id)`가 막는다. 덤프 경로는 첫 줄만 쓰고 뒤 줄은 stderr에 알리고 버린다.
8. **디코드되지 않는 줄.** wikipgeul는 `load_error`로 남기고 개체를 만들지 않는다. 덤프 경로도 그 개체를 쓰지 않는다.
   - 정렬 단계: `id`를 못 읽거나 ID가 나뉘지 않으면 "not read"로 세고 줄 번호를 표준 출력에 적는다.
   - 변환 단계: 디코드 실패, 또는 디코드한 ID가 정렬 때 읽은 첫 `"id"`와 다르면(대소문자만 다른 키·중복 `id`) 그 개체를 빼고 stderr에 알린다. 보고서 줄도 없다. DB 경로도 그 개체를 나열하지 않는다.
   - 엄격 모드에서는 이상점 `decode`(`덤프 줄 해석 실패`, 데이터 이상)로 멈춘다. `-allow decode:덤프 줄 해석 실패`로 넘긴다.
   - 사전 학습 표본도 DB 경로와 같다(`sampleJob`). 목록 없이 고를 때는 그런 줄을 세지 않고 다음 Q를 쓴다(DB는 그 개체를 나열하지 않는다). 목록에서 고를 때는 없는 개체로 남겨 건너뛴다.
9. **보고서 첫 줄.** `-schema … -dump-id …` 대신 `-dump <파일 이름> sha256:<hex>`이다. 나머지 열은 같다.

## 순서

gogeul `pack.Stream.Put`은 키(Issuer, 문서 ID, slot)가 **엄격히 오름차순**이어야 한다. 정렬 런 + 합치기는 스트림 안에서 H 키(내용 해시)에만 쓰인다. 위키데이터 덤프는 Q 번호 순이 아니고 P가 섞여 있으므로 순서를 맞춰야 한다. wikigeul이 두 단계로 처리한다.

1. **정렬(`dumpfile.Sort`, `-sort-dir`).**
   - 덤프를 한 번 읽어 이번 실행의 개체 줄만 남긴다(`keepEntity`). 기본은 모든 P(번호 2^56−1 이하)와 Q<from>…Q<to>, `-ids`면 목록의 개체다.
   - 개체 JSON이 `-sort-mem` MiB만큼 모이면 (종류, 번호, 줄 번호) 순으로 정렬해 zstd 런 파일 하나로 쓴다(가장 빠른 수준, 1 MiB 창).
   - 다 읽으면 런을 64개씩 합쳐 64개 이하로 줄인다.
   - 디스크: 남긴 JSON을 zstd로 압축한 크기 정도다(전체 변환이면 원본 JSON의 대략 1/8~1/10). 합치기 중에는 합치는 런만큼 더 쓴다.
2. **변환.**
   - 남은 런을 k-way로 합치면(`dumpfile.Merger`) 문서 키 순서의 개체 줄이 된다. P(0x50…)가 모두 Q(0x51…) 앞이고, 그다음은 번호 순이다.
   - 이 순서로 작업 번호를 매겨 작업자에게 보내고, 기존 작성자가 번호 순으로 스트림에 넣는다.
   - 메모리는 런마다 디코더 하나와 레코드 하나, 그리고 작업 채널·대기 결과다.

zstd 사전 표본은 DB 경로와 같은 개체다(`dumpSampleJobs`). 목록이 없으면 정렬된 덤프의 첫 `-samples`개 Q이고, `-ids`면 목록의 첫 `-samples`개 Q(덤프에 없는 것은 건너뜀)다.

## 재개

| 단계 | 체크포인트 | 재개 |
|---|---|---|
| 정렬 | 런을 쓸 때마다, 또는 남긴 것이 적어도 덤프 텍스트 16 GiB마다 `state.json`을 원자 저장한다(임시 파일 → fsync → rename → 디렉터리 fsync). 담는 것: 읽은 줄 수, 텍스트 바이트 오프셋, 런 목록, 평문이면 그 오프셋까지의 SHA-256 상태 | 같은 명령을 다시 실행한다(`-resume`과 무관하게 `-sort-dir`의 상태를 이어 간다). 덤프 파일(이름·크기·수정 시각·SHA-256)이나 범위·목록이 다르면 거부한다 |
| 합치기 | 합치기 한 번마다(새 런은 임시 이름 → rename, 상태 저장 뒤 입력 런 삭제) | 중단된 합치기를 같은 이름으로 다시 한다 |
| 변환 | DB 경로와 같다: `-checkpoint` 문서마다 pack 스트림·보고서·진행을 한 원자 저장에 | `-resume`. 런을 처음부터 다시 합치되 체크포인트 개체까지는 변환하지 않고 넘긴다(런 압축 해제 비용만 든다). 정렬 중에 멈춘 실행에 `-resume`을 줘도 스트림이 아직 없으면 새로 시작한다 |

압축 형식별로 정렬 단계를 어디서 재개할 수 있는지:

- **평문 `.json`:** 정확한 바이트 오프셋으로 seek하고 SHA-256 상태(`encoding.BinaryMarshaler`)를 되살린다. 다시 읽는 것이 없다.
- **`.json.gz`:** 처음부터 다시 압축을 풀고(파일 SHA-256도 다시 계산) 오프셋까지의 텍스트를 버린다. deflate 스트림 중간에서 시작하려면 비트 오프셋과 직전 32 KiB 창이 필요한데(zran 방식) Go의 gzip 구현이 내부 상태를 내놓지 않는다. 멀티 멤버 파일은 멤버 경계에서 시작할 수 있지만 구현하지 않았다(위키미디어 파일은 단일 멤버로 본다).
- **`.json.bz2`:** gz와 같다. 처음부터 다시 풀어 버린다. bzip2 블록은 비트 단위 경계라 바이트 오프셋으로 들어갈 수 없다. lbzip2·bzip2는 중간에서 시작하는 옵션이 없고, Go `compress/bzip2`는 seek가 없고 너무 느리다. 다시 푸는 비용은 lbzip2 병렬 속도의 해제 시간뿐이다(런은 다시 쓰지 않는다).

## 플래그 관계

| 플래그 | 뜻 |
|---|---|
| `-dump PATH` | 읽을 덤프 파일(필수). DB에서 변환하려면 wikipgeul `export-geul`을 쓴다 |
| `-from`, `-to`, `-ids` | DB 경로와 같은 개체 집합(정렬 필터). 정렬 상태의 정체에도 들어간다. `-ids`에 있으나 덤프에 없는 개체는 DB 경로처럼 보고서 `missing`(엄격 모드면 이상점) |
| `-workers` | 변환 goroutine 수. 정렬 때 lbzip2 스레드·런 zstd 인코더 스레드 수. 기본은 CPU의 70%(wikipgeul `export-geul`은 DB 백엔드를 세어 그 절반) |
| `-checkpoint` | 변환 단계만. 정렬은 런마다 체크포인트 |
| `-resume` | 변환 단계만. 정렬은 늘 이어 간다 |
| `-report`, `-strict`, `-anomaly`, `-allow` | DB 경로와 같다. 이상점 종류 `decode`가 늘었다 |
| `-recorded-at`, `-samples`, `-langreg-root`, `-langreg-version`, `-commit`, `-lib-commit` | 같다. DB 경로와 같은 문서를 원하면 `-recorded-at`을 같게 준다 |
| `-dump-url`, `-dump-sha256`, `-dump-time` | 문서 HEAD의 덤프 출처(위 대응표). 비우면 파일에서 정한다 |
| `-sort-dir` | 정렬 상태와 런. 기본 `wikigeul-dumpsort`(현재 디렉터리 기준). 아카이브(`-root`) 밖에 둔다 |
| `-sort-mem` | 런 하나에 모을 개체 JSON MiB(기본 1024). 메모리 사용은 대략 이 값과 줄 버퍼 정도다. 작게 주면 런이 많아져 합치기 단계가 늘어난다 |

체크포인트 정체(`pipeline.PackProgress.Source`)에 덤프 파일 이름과 SHA-256이 들어간다. DB 경로의 정체는 그 필드가 비어 있어 이전 체크포인트와 그대로 맞는다. wikigeul에는 `-schema`·`-dump-id`가 없고, 정체의 `Schema`·`Dump`는 분리 전 체크포인트와 맞도록 기본값(`wikidata`, 1)으로 둔다.

## 검증

단위 테스트는 작은 합성 덤프로 돈다. 대상은 `pkg/source`(jsonb·numeric 출력, `DecodeDumpEntity`), `internal/dumpfile`(평문·gz·bz2 정렬, 재개, 합치기, SHA-256 검증), `pkg/pipeline`(pack 쓰기, 엄격 모드, 보고서), `cmd/wikigeul`(전체 실행, 표본, 작업 생성)이다. DB 경로와 같은 바이트인지는 실제 덤프와 적재본으로만 확인할 수 있다. DB 경로의 전체 변환이 끝난 뒤 다음 순서로 확인한다.

1. **실행 조건 확인.** DB 경로 실행 출력의 `reference snak order: … true`와 wikipgeul `build-reference-order` 보고의 `content` 수를 확인한다. DB `dump` 행에서 `url`, `sha256`, `floor(extract(epoch FROM downloaded_at))`을 읽어 둔다.
2. **표본 비교.** `-ids` 목록을 하나 만든다.
   - 무작위 Q 1만 개
   - 경계 사례: 속성 전부, `\u0000` 값이 있는 개체, 출처가 많고 순서가 바뀌는 개체, 좌표 정밀도 `1.0e-5`류, quantity, 배지 있는 sitelink, 원문 보존 사유별 개체(이전 실행 보고서의 `raw` 열에서 고른다)

   두 경로를 같은 `-commit`(wikigeul 커밋)·`-lib-commit`·`-langreg-version`·`-recorded-at`·목록으로 각각 다른 `-root`에 돌린다. 덤프 경로에는 1의 값을 `-dump-url`·`-dump-sha256`·`-dump-time`으로 준다.

   ```sh
   wikipgeul export-geul -ids sample.txt -root /tmp/pg -recorded-at T … -report pg.tsv
   wikigeul -ids sample.txt -root /tmp/dump -recorded-at T -dump <DUMP_PATH> -sort-dir /tmp/sort … -report dump.tsv
   ```

3. **비교 기준.**
   - 두 pack 세트에서 문서마다 버전 1의 내용 해시(H 키, `h/<hash>.geul`)와 버전 목록 해시가 모두 같아야 한다. 같으면 문서 바이트가 같다.
   - 보고서는 첫 줄을 빼고 `diff`가 비어야 한다.
   - 다르면 그 개체를 레코드 단위로 비교한다. DB 경로는 wikipgeul `export-geul -out dir`로 풀고, 덤프 경로는 디렉터리 모드가 없으므로 pack에서 꺼낸다. 그 줄의 `source.Entity`를 두 경로에서 뽑아 비교하는 Go 테스트를 붙여도 된다.
4. **전체 비교(선택).** 같은 조건으로 덤프 경로 전체 변환을 돌려 두 세트의 P 색인(문서 → 버전 1 해시)을 전수 비교한다. 다른 문서는 위 4번 항목(출처 hash 내용 차이) 같은 알려진 예외인지 분류한다.
