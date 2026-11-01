//ff:func feature=memguard type=helper control=sequence
//ff:what TestMachineStop — Danger with no worker, memory given back and a reason naming the limit
package memguard

import (
	"strings"
	"testing"
)

func TestMachineStop(t *testing.T) {
	m := fixtureMachine(4)
	d := decision{Press: pressure{Name: "cgroup", Free: 1, Size: 4}}
	m.stop(&d)
	if m.level != Danger || m.active != 0 || !d.Reclaim || !strings.Contains(d.Reason, "cgroup free 25.0%") {
		t.Fatalf("%+v", d)
	}
}
