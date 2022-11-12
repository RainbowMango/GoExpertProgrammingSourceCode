package examples

import (
	"reflect"
	"sort"
	"testing"
)

func TestPreviousUsage(t *testing.T) {
	s := []int{3, 5, 2}
	sort.Slice(s, func(i, j int) bool {
		if s[i] < s[j] {
			return true
		}
		return false
	})

	if !reflect.DeepEqual([]int{2, 3, 5}, s) {
		t.Fatalf("unexpected, got: %v", s)
	}
}

func TestGenericUsage(t *testing.T) {
	s := []int{3, 5, 2}
	OrderedSlice(s)

	if !reflect.DeepEqual([]int{2, 3, 5}, s) {
		t.Fatalf("unexpected, got: %v", s)
	}
}
