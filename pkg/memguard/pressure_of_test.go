//ff:func feature=memguard type=helper control=sequence
//ff:what TestPressureOf — the system alone; a tighter cgroup; the cap close to its floors; machine memory plentiful does not hide them
package memguard

import "testing"

func TestPressureOf(t *testing.T) {
	l := Limits{}.withDefaults()
	s := Sample{RSS: 500, Total: 1000, Avail: 500}
	if p := pressureOf(l, s); p.Name != "system" || p.Caution != 200 || p.Danger != 350 || p.Recover != 150 {
		t.Fatalf("%+v", p)
	}
	s.CgroupLimit, s.CgroupUsed = 1000, 900
	if p := pressureOf(l, s); p.Name != "cgroup" || p.Caution != -200 || p.Free != 100 || p.Size != 1000 {
		t.Fatalf("%+v", p)
	}
	s.CgroupLimit = 0
	l.Max = 600
	if p := pressureOf(l, s); p.Name != "max" || p.Free != 100 || p.Caution != 100-120 || p.Danger != 100-30 || p.Recover != 100-180 {
		t.Fatalf("%+v", p)
	}
}
