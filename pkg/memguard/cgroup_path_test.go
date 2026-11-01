//ff:func feature=memguard type=loader control=sequence
//ff:what TestCgroupPath — the v2 line; the v1 memory controller's, also among others and on a hybrid system; short lines ignored
package memguard

import "testing"

func TestCgroupPath(t *testing.T) {
	if rel, v2 := cgroupPath([]byte("0::/system.slice/x.service\n\nbad\n")); rel != "/system.slice/x.service" || !v2 {
		t.Fatalf("%q %v", rel, v2)
	}
	if rel, v2 := cgroupPath([]byte("9:cpu,memory:/y\n")); rel != "/y" || v2 {
		t.Fatalf("%q %v", rel, v2)
	}
	if rel, v2 := cgroupPath([]byte("0::/v2\n3:memory:/v1\n")); rel != "/v1" || v2 {
		t.Fatalf("%q %v", rel, v2)
	}
}
