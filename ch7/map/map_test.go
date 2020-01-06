package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m1 := map[int]int{1:1, 2:4, 3:9}
	t.Log(m1[2])
	t.Logf("len of m1 = %d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len of m2 = %d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len of m3 = %d", len(m3))
}

func TestMapExistingKey(t *testing.T) {
	m := map[int]int{}
	t.Log(m[1])
	m[2] = 0
	t.Log(m[2])
	//m[3] = 0
	if v, ok := m[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestMapTravel(t *testing.T) {
	m := map[int]int{1:1, 2:4, 3:9}
	for k, v := range m {
		t.Log(k, v)
	}
}