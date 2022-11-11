package examples

import "testing"

func TestMapKeys(t *testing.T) {
	m := map[int]int{1: 2, 2: 4}
	k := MapKeys(m)
	if len(k) != 2 {
		t.Fatalf("Expect two eletems, bug got: %d", len(k))
	}
}
