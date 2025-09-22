package genomicrange

import "testing"

func TestGenomicRange(t *testing.T) {

	tests := []struct {
		S        string
		P, Q     []int
		expected []int
	}{
		{ 
			"CAGCCTA", []int{2, 5, 0, 1, 3, 6, 3}, []int{4, 5, 6, 5, 6, 6, 5}, []int{2, 4, 1, 1, 1, 1, 2},
		},
		{
			"AC", []int{0, 0, 1}, []int{0, 1, 1}, []int{1, 1, 2},
		},
	}

	for _, test := range tests {
		result := Solution(test.S, test.P, test.Q)
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("GenomicRange(%s) = %#v; expected %#v", test.S, result, test.expected)
			}
		}
	}
}
