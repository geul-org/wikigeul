//ff:func feature=pipeline type=helper control=sequence
//ff:what TestAnomalyErrorError — the message names the count, entity and report path
package pipeline

import "testing"

func TestAnomalyErrorError(t *testing.T) {
	if got := (AnomalyError{ID: "Q7", Count: 2, Path: "a.tsv"}).Error(); got != "wikigeul -strict: 2 anomalies at Q7, stopped after a checkpoint; report a.tsv" {
		t.Fatal(got)
	}
}
