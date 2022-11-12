package examples

import (
	"testing"
)

func TestSet(t *testing.T) {
	s := MakeSet[int]()

	tests := []struct {
		name    string
		inputs  []int
		outputs []int
	}{
		{
			name:    "thread1 to crud",
			inputs:  []int{1, 2, 3},
			outputs: []int{1, 2, 3},
		},
		{
			name:    "thread2 to crud",
			inputs:  []int{1, 2, 3},
			outputs: []int{1, 2, 3},
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			for i := range tc.inputs {
				s.Add(tc.inputs[i])
				if !s.Contains(tc.inputs[i]) {
					t.Fatalf("unexpected")
				}
			}
		})
	}

}
