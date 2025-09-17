package frogriver

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		X        int
		A        []int
		expected int
	}{
		{5, []int{1, 3, 1, 4, 2, 3, 5, 4}, 6},
		{5, []int{1, 2, 3, 4, 5}, 4},
		{5, []int{5, 4, 3, 2, 1}, 4},
		{0, []int{}, 0},
		{1, []int{1}, 0},
		{10, []int{1,2,3,3,3,1,2,1,2,1,2,1,2,2,3,4,1,1,1,5,6,7,8,9,1,1,10,1,1,1,1,1,1,1,1,1,1}, 26},
	}

	for _, test := range tests {
		result := Solution(test.X, test.A)
		if result != test.expected {
			t.Errorf("FrogRiver(%d, %v) = %d; expected %d", test.X, test.A, result, test.expected)
		}
	}
}
