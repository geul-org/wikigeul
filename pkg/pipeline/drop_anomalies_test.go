//ff:func feature=pipeline type=command control=sequence
//ff:what TestDropAnomalies — one decode anomaly with the error as its value, none when -allow accepts it
package pipeline

import "testing"

func TestDropAnomalies(t *testing.T) {
	got := dropAnomalies("Q7", "why", nil)
	if len(got) != 1 || got[0] != (anomaly{ID: "Q7", Kind: "decode", Value: "why", Reason: dropReason, Class: "데이터 이상"}) {
		t.Fatalf("%+v", got)
	}
	if got := dropAnomalies("Q7", "why", allowSet([]string{"decode:" + dropReason})); got != nil {
		t.Fatalf("allowed: %+v", got)
	}
}
