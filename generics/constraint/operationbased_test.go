package constraint

import "testing"

func TestGreaterThan(t *testing.T) {
	var aInt, bInt = 1, 2
	if GreaterThan(aInt, bInt) {
		t.Fatalf("aInt(%d) should not greater than bInt(%d)", aInt, bInt)
	}

	var aStr, bStr = "aaa", "bbb"
	if GreaterThan(aStr, bStr) {
		t.Fatalf("aStr(%s) should not greater than bStr(%s)", aStr, bStr)
	}
}
