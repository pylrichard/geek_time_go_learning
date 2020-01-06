package map_ext

import "testing"

func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	t.Log(m[1](2), m[2](2))
}

func TestMapForSet(t *testing.T) {
	set := map[int]bool{}
	set[1] = true
	n := 3
	if set[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	set[3] = true
	t.Log(len(set))
	delete(set, 1)
	n = 1
	if set[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}