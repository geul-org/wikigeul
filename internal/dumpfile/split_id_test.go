//ff:func feature=dumpfile type=loader control=iteration dimension=1
//ff:what TestSplitID — letter and number of Q and P IDs; leading zero, sign, no number or overflow is false
package dumpfile

import "testing"

func TestSplitID(t *testing.T) {
	if k, n, ok := SplitID("Q42"); !ok || k != 'Q' || n != 42 {
		t.Fatalf("Q42: %c %d %v", k, n, ok)
	}
	for _, bad := range []string{"", "Q", "Q0", "Q01", "Q-1", "Q+1", "Qx", "Q99999999999999999999"} {
		if _, _, ok := SplitID(bad); ok {
			t.Errorf("%q: want false", bad)
		}
	}
}
